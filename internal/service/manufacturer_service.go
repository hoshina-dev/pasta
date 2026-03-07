package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/hoshina-dev/pasta/internal/repository"
)

type ManufacturerService struct {
	repo repository.ManufacturerRepository
}

func NewManufacturerService(repo repository.ManufacturerRepository) *ManufacturerService {
	return &ManufacturerService{repo: repo}
}

func (s *ManufacturerService) Create(ctx context.Context, input model.CreateManufacturerInput) (*model.Manufacturer, error) {
	m := &model.Manufacturer{
		Name:            input.Name,
		CountryOfOrigin: input.CountryOfOrigin,
	}
	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *ManufacturerService) GetByID(ctx context.Context, id uuid.UUID) (*model.Manufacturer, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ManufacturerService) GetAll(ctx context.Context, limit, offset int) ([]model.Manufacturer, error) {
	return s.repo.GetAll(ctx)
}

func (s *ManufacturerService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
