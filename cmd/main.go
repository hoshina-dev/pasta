package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	appConfig "github.com/hoshina-dev/pasta/internal/config"
	"github.com/hoshina-dev/pasta/internal/graphql"
	"github.com/hoshina-dev/pasta/internal/infra/postgres"
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

	partRepo := repository.NewPartRepository(db)
	manufacturerRepo := repository.NewManufacturerRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	partSvc := service.NewPartService(partRepo, manufacturerRepo, categoryRepo)
	manufacturerSvc := service.NewManufacturerService(manufacturerRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	storageSvc := service.NewStorageService(s3Storage)

	resolver := graphql.NewResolver(partSvc, manufacturerSvc, categorySvc, storageSvc)

	app := server.New(resolver, cfg.CORSOrigins)

	log.Printf("starting server on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
