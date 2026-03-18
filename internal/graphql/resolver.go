package graphql

import "github.com/hoshina-dev/pasta/internal/service"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	partService         *service.PartService
	manufacturerService *service.ManufacturerService
	categoryService     *service.CategoryService
	storageService      *service.StorageService
	optimizationService *service.OptimizationService
}

func NewResolver(partSvc *service.PartService, manufacturerSvc *service.ManufacturerService, categorySvc *service.CategoryService, storageSvc *service.StorageService, optimizationSvc *service.OptimizationService) *Resolver {
	return &Resolver{
		partService:         partSvc,
		manufacturerService: manufacturerSvc,
		categoryService:     categorySvc,
		storageService:      storageSvc,
		optimizationService: optimizationSvc,
	}
}
