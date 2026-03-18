package model

import (
	"time"

	"github.com/google/uuid"
)

type OptimizationJobLog struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	JobID uuid.UUID `gorm:"type:uuid;not null;index"`

	SourceURL string  `gorm:"type:text;not null"`
	DestURL   string  `gorm:"type:text;not null"`
	SourceKey *string `gorm:"type:text"`
	DestKey   *string `gorm:"type:text"`

	DracoCompressionLevel     *int `gorm:"type:int"`
	DracoPositionQuantization *int `gorm:"type:int"`
	DracoTexcoordQuantization *int `gorm:"type:int"`
	DracoNormalQuantization   *int `gorm:"type:int"`
	DracoGenericQuantization  *int `gorm:"type:int"`

	Status       string  `gorm:"type:text;not null;index"`
	ExitCode     *int    `gorm:"type:int"`
	ErrorMessage *string `gorm:"type:text"`

	SourceFileSize    *int64   `gorm:"type:bigint"`
	ProcessedFileSize *int64   `gorm:"type:bigint"`
	CompressionRatio  *float64 `gorm:"type:decimal(5,2)"`

	StartedAt       *time.Time `gorm:"type:timestamp with time zone"`
	CompletedAt     *time.Time `gorm:"type:timestamp with time zone;index"`
	DurationSeconds *int       `gorm:"type:int;index"`

	JobLogs string `gorm:"type:text"`

	WebhookReceivedAt time.Time `gorm:"type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`
	CreatedAt         time.Time `gorm:"type:timestamp with time zone;not null;autoCreateTime;index"`
	UpdatedAt         time.Time `gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
}
