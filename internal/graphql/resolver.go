package graphql

import "github.com/hoshina-dev/pasta/internal/service"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	pastaService        *service.PastaService
	manufacturerService *service.ManufacturerService
	categoryService     *service.CategoryService
}

func NewResolver(partSvc *service.PastaService, manufacturerSvc *service.ManufacturerService, categorySvc *service.CategoryService) *Resolver {
	return &Resolver{pastaService: partSvc, manufacturerService: manufacturerSvc, categoryService: categorySvc}
}
