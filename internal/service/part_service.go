package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/hoshina-dev/pasta/internal/repository"
)

type PartService struct {
	partRepo         repository.PartRepository
	manufacturerRepo repository.ManufacturerRepository
	categoryRepo     repository.CategoryRepository
}

func NewPartService(partRepo repository.PartRepository, manufacturerRepo repository.ManufacturerRepository, categoryRepo repository.CategoryRepository) *PartService {
	return &PartService{partRepo: partRepo, manufacturerRepo: manufacturerRepo, categoryRepo: categoryRepo}
}

func (s *PartService) GetAll(ctx context.Context) ([]model.Part, error) {
	return s.partRepo.GetAll(ctx)
}

func (s *PartService) GetByID(ctx context.Context, id uuid.UUID) (*model.Part, error) {
	return s.partRepo.GetByID(ctx, id)
}

func (s *PartService) Search(ctx context.Context, name string) ([]model.Part, error) {
	return s.partRepo.Search(ctx, name)
}

func (s *PartService) Create(ctx context.Context, input model.CreatePartInput) (*model.Part, error) {
	// Verify manufacturer exists
	m, err := s.manufacturerRepo.GetByID(ctx, input.ManufacturerID)
	if err != nil || m == nil {
		return nil, fmt.Errorf("manufacturer with id %s not found", input.ManufacturerID)
	}

	// Verify categories exists
	categories, err := s.categoryRepo.GetByIDs(ctx, input.CategoryIDs)
	if err != nil {
		return nil, err
	}
	if len(categories) != len(input.CategoryIDs) {
		return nil, fmt.Errorf("one or more categories not found")
	}

	part := input.ToModel()
	part.Manufacturer = *m
	part.Categories = categories
	if err := s.partRepo.Create(ctx, part); err != nil {
		return nil, err
	}

	return part, nil
}

func (s *PartService) Update(ctx context.Context, id uuid.UUID, input model.UpdatePartInput) (*model.Part, error) {
	// Verify part exists
	part, err := s.partRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Verify categories exists
	categories, err := s.categoryRepo.GetByIDs(ctx, input.CategoryIDs)
	if err != nil {
		return nil, err
	}
	if len(categories) != len(input.CategoryIDs) {
		return nil, fmt.Errorf("one or more categories not found")
	}

	model.ApplyUpdatePartInput(part, input)
	part.Categories = categories
	if err := s.partRepo.Update(ctx, part); err != nil {
		return nil, err
	}
	return part, nil
}

func (s *PartService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.partRepo.Delete(ctx, id)
}
