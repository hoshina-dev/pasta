package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/hoshina-dev/pasta/internal/repository"
)

type WebhookHandler struct {
	part3DModelRepo repository.Part3DModelRepository
	jobLogRepo      repository.OptimizationJobLogRepository
	s3BaseURL       string
}

func NewWebhookHandler(
	part3DModelRepo repository.Part3DModelRepository,
	jobLogRepo repository.OptimizationJobLogRepository,
	s3BaseURL string,
) *WebhookHandler {
	return &WebhookHandler{
		part3DModelRepo: part3DModelRepo,
		jobLogRepo:      jobLogRepo,
		s3BaseURL:       s3BaseURL,
	}
}

type OptimizationWebhookPayload struct {
	UUID                      string     `json:"uuid"`
	Status                    string     `json:"status"`
	ExitCode                  int        `json:"exit_code"`
	Logs                      string     `json:"logs"`
	Timestamp                 time.Time  `json:"timestamp"`
	SourceURL                 string     `json:"source_url,omitempty"`
	DestURL                   string     `json:"dest_url,omitempty"`
	SourceFileSize            *int64     `json:"source_file_size,omitempty"`
	ProcessedFileSize         *int64     `json:"processed_file_size,omitempty"`
	DracoCompressionLevel     *int       `json:"draco_compression_level,omitempty"`
	DracoPositionQuantization *int       `json:"draco_position_quantization,omitempty"`
	DracoTexcoordQuantization *int       `json:"draco_texcoord_quantization,omitempty"`
	DracoNormalQuantization   *int       `json:"draco_normal_quantization,omitempty"`
	DracoGenericQuantization  *int       `json:"draco_generic_quantization,omitempty"`
	StartedAt                 *time.Time `json:"started_at,omitempty"`
	CompletedAt               *time.Time `json:"completed_at,omitempty"`
	DurationSeconds           *int       `json:"duration_seconds,omitempty"`
}

func (h *WebhookHandler) HandleOptimizationCallback(c *fiber.Ctx) error {
	ctx := c.UserContext()
	if ctx == nil {
		ctx = context.Background()
	}

	var payload OptimizationWebhookPayload
	if err := c.BodyParser(&payload); err != nil {
		log.Printf("failed to parse webhook payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	log.Printf("received optimization webhook: uuid=%s, status=%s, exit_code=%d",
		payload.UUID, payload.Status, payload.ExitCode)

	jobID, err := uuid.Parse(payload.UUID)
	if err != nil {
		log.Printf("invalid UUID in webhook: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid UUID",
		})
	}

	model3D, err := h.part3DModelRepo.GetByID(ctx, jobID)
	if err != nil {
		log.Printf("failed to get 3D model: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get 3D model",
		})
	}
	if model3D == nil {
		log.Printf("3D model not found: %s", jobID)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "3D model not found",
		})
	}

	var status model.Part3DModelStatus
	var processedURL *string

	if payload.Status == "success" && payload.ExitCode == 0 {
		status = model.Part3DModelStatusReady
		if model3D.ProcessedKey != nil {
			url := fmt.Sprintf("%s/%s", h.s3BaseURL, *model3D.ProcessedKey)
			processedURL = &url
		}
	} else {
		status = model.Part3DModelStatusFailed
		log.Printf("optimization job failed - logs:\n%s", payload.Logs)
	}

	err = h.part3DModelRepo.UpdateStatus(ctx, jobID, string(status), processedURL)
	if err != nil {
		log.Printf("failed to update 3D model status: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update status",
		})
	}

	if err := h.logJobExecution(ctx, jobID, model3D, payload); err != nil {
		log.Printf("failed to log job execution: %v", err)
	}

	log.Printf("successfully updated 3D model status: uuid=%s, status=%s", payload.UUID, status)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "webhook processed",
	})
}

func (h *WebhookHandler) logJobExecution(ctx context.Context, jobID uuid.UUID, model3D *model.Part3DModel, payload OptimizationWebhookPayload) error {
	var compressionRatio *float64
	if payload.SourceFileSize != nil && payload.ProcessedFileSize != nil && *payload.SourceFileSize > 0 {
		ratio := (1.0 - float64(*payload.ProcessedFileSize)/float64(*payload.SourceFileSize)) * 100.0
		compressionRatio = &ratio
	}

	var errorMessage *string
	if payload.Status != "success" || payload.ExitCode != 0 {
		msg := fmt.Sprintf("Job failed with exit code %d", payload.ExitCode)
		errorMessage = &msg
	}

	var sourceKey, destKey *string
	if model3D.RawURL != "" {
		sourceKey = &model3D.RawURL
	}
	if model3D.ProcessedKey != nil {
		destKey = model3D.ProcessedKey
	}

	// Resolve source and destination URLs, falling back to model data if needed.
	sourceURL := payload.SourceURL
	if sourceURL == "" && model3D.RawURL != "" {
		sourceURL = model3D.RawURL
	}

	destURL := payload.DestURL
	if destURL == "" && model3D.ProcessedKey != nil {
		// Construct destination URL from the S3 base URL and the processed key.
		destURL = h.s3BaseURL + "/" + *model3D.ProcessedKey
	}

	// Ensure we have non-empty URLs before creating the log entry to satisfy DB constraints.
	if sourceURL == "" || destURL == "" {
		return fmt.Errorf("missing source or destination URL for optimization job %s", jobID.String())
	}

	jobLog := &model.OptimizationJobLog{
		JobID:                     jobID,
		SourceURL:                 sourceURL,
		DestURL:                   destURL,
		SourceKey:                 sourceKey,
		DestKey:                   destKey,
		DracoCompressionLevel:     payload.DracoCompressionLevel,
		DracoPositionQuantization: payload.DracoPositionQuantization,
		DracoTexcoordQuantization: payload.DracoTexcoordQuantization,
		DracoNormalQuantization:   payload.DracoNormalQuantization,
		DracoGenericQuantization:  payload.DracoGenericQuantization,
		Status:                    payload.Status,
		ExitCode:                  &payload.ExitCode,
		ErrorMessage:              errorMessage,
		SourceFileSize:            payload.SourceFileSize,
		ProcessedFileSize:         payload.ProcessedFileSize,
		CompressionRatio:          compressionRatio,
		StartedAt:                 payload.StartedAt,
		CompletedAt:               payload.CompletedAt,
		DurationSeconds:           payload.DurationSeconds,
		JobLogs:                   payload.Logs,
		WebhookReceivedAt:         time.Now(),
	}

	return h.jobLogRepo.Create(ctx, jobLog)
}
