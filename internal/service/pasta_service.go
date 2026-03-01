package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/hoshina-dev/pasta/internal/repository"
)

type PastaService struct {
	partRepo         repository.PastaRepository
	manufacturerRepo repository.ManufacturerRepository
	categoryRepo     repository.CategoryRepository
}

func NewPastaService(partRepo repository.PastaRepository, manufacturerRepo repository.ManufacturerRepository, categoryRepo repository.CategoryRepository) *PastaService {
	return &PastaService{partRepo: partRepo, manufacturerRepo: manufacturerRepo, categoryRepo: categoryRepo}
}

func (s *PastaService) GetAll(ctx context.Context) ([]model.Pasta, error) {
	return s.partRepo.GetAll(ctx)
}

func (s *PastaService) GetByID(ctx context.Context, id uuid.UUID) (*model.Pasta, error) {
	return s.partRepo.GetByID(ctx, id)
}

func (s *PastaService) Search(ctx context.Context, name string) ([]model.Pasta, error) {
	return s.partRepo.Search(ctx, name)
}

func (s *PastaService) Create(ctx context.Context, input model.CreatePastaInput) (*model.Pasta, error) {
	// Verify manufacturer exists
	m, err := s.manufacturerRepo.GetByID(ctx, input.ManufacturerID)
	if err != nil || m == nil {
		return nil, fmt.Errorf("manufacturer with id %s not found", input.ManufacturerID)
	}

	// Verify categories exists
	categories := make([]model.Category, len(input.CategoryIDs))
	for i, cid := range input.CategoryIDs {
		c, err := s.categoryRepo.GetByID(ctx, cid)
		if err != nil || c == nil {
			return nil, fmt.Errorf("category with id %s not found", cid)
		}
		categories[i] = *c
	}

	pasta := input.ToModel()
	pasta.Manufacturer = *m
	pasta.Categories = categories
	if err := s.partRepo.Create(ctx, pasta); err != nil {
		return nil, err
	}

	return pasta, nil
}

func (s *PastaService) Update(ctx context.Context, id uuid.UUID, input model.UpdatePastaInput) (*model.Pasta, error) {
	pasta, err := s.partRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	model.ApplyUpdatePartInput(pasta, input)
	if err := s.partRepo.Update(ctx, pasta); err != nil {
		return nil, err
	}
	return pasta, nil
}

func (s *PastaService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.partRepo.Delete(ctx, id)
}
