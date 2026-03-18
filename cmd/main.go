package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	appConfig "github.com/hoshina-dev/pasta/internal/config"
	"github.com/hoshina-dev/pasta/internal/graphql"
	"github.com/hoshina-dev/pasta/internal/handler"
	"github.com/hoshina-dev/pasta/internal/infra/postgres"
	"github.com/hoshina-dev/pasta/internal/infra/rabbitmq"
	storage "github.com/hoshina-dev/pasta/internal/infra/s3"
	"github.com/hoshina-dev/pasta/internal/repository"
	"github.com/hoshina-dev/pasta/internal/server"
	"github.com/hoshina-dev/pasta/internal/service"
)

func main() {
	cfg := appConfig.Load()

	db, err := postgres.Connect(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	ctx := context.Background()
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.S3Region))
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	s3Client := s3.NewFromConfig(awsCfg)
	s3Storage := storage.NewS3StorageService(s3Client, cfg.S3Bucket, cfg.S3BaseURL)

	rabbitConn, rabbitCh, err := rabbitmq.Connect(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	if err := rabbitmq.DeclareExchange(rabbitCh, cfg.RabbitMQExchange); err != nil {
		log.Fatalf("failed to declare RabbitMQ exchange: %v", err)
	}

	rabbitPublisher := rabbitmq.NewRabbitPublisher(rabbitCh)

	partRepo := repository.NewPartRepository(db)
	manufacturerRepo := repository.NewManufacturerRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	part3DModelRepo := repository.NewPart3DModelRepository(db)
	jobLogRepo := repository.NewOptimizationJobLogRepository(db)

	webhookURL := cfg.OptimizationWebhookURL
	if webhookURL == "" {
		webhookURL = fmt.Sprintf("http://localhost:%s/webhook/optimization", cfg.Port)
		log.Printf("OPTIMIZATION_WEBHOOK_URL not set, using default: %s", webhookURL)
	}

	partSvc := service.NewPartService(partRepo, manufacturerRepo, categoryRepo)
	manufacturerSvc := service.NewManufacturerService(manufacturerRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	storageSvc := service.NewStorageService(s3Storage)
	optimizationSvc := service.NewOptimizationService(
		s3Storage,
		rabbitPublisher,
		part3DModelRepo,
		webhookURL,
		cfg.RabbitMQExchange,
		cfg.RabbitMQRoutingKey,
	)

	webhookHandler := handler.NewWebhookHandler(part3DModelRepo, jobLogRepo, cfg.S3BaseURL)
	resolver := graphql.NewResolver(partSvc, manufacturerSvc, categorySvc, storageSvc, optimizationSvc)

	app := server.New(resolver, webhookHandler, cfg.CORSOrigins)

	log.Printf("starting server on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
