package main

import (
	"log"

	"github.com/hoshina-dev/pasta/internal/config"
	"github.com/hoshina-dev/pasta/internal/graphql"
	"github.com/hoshina-dev/pasta/internal/infra/postgres"
	"github.com/hoshina-dev/pasta/internal/repository"
	"github.com/hoshina-dev/pasta/internal/server"
	"github.com/hoshina-dev/pasta/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := postgres.Connect(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	pastaRepo := repository.NewPastaRepository(db)
	manufacturerRepo := repository.NewManufacturerRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	pastaSvc := service.NewPastaService(pastaRepo, manufacturerRepo, categoryRepo)
	manufacturerSvc := service.NewManufacturerService(manufacturerRepo)
	categorySvc := service.NewCategoryService(categoryRepo)

	resolver := graphql.NewResolver(pastaSvc, manufacturerSvc, categorySvc)

	app := server.New(resolver, cfg.CORSOrigins)

	log.Printf("starting server on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
