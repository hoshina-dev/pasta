package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/hoshina-dev/pasta/internal/repository"
)

type PastaService struct {
	repo repository.PastaRepository
}

func NewPastaService(repo repository.PastaRepository) *PastaService {
	return &PastaService{repo: repo}
}

func (s *PastaService) GetAll(ctx context.Context) ([]model.Pasta, error) {
	return s.repo.GetAll(ctx)
}

func (s *PastaService) GetByID(ctx context.Context, id uuid.UUID) (*model.Pasta, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PastaService) Search(ctx context.Context, name string) ([]model.Pasta, error) {
	return s.repo.Search(ctx, name)
}

func (s *PastaService) Create(ctx context.Context, name, description string, price float64) (*model.Pasta, error) {
	pasta := &model.Pasta{
		Name:        name,
		Description: &description,
	}
	if err := s.repo.Create(ctx, pasta); err != nil {
		return nil, err
	}
	return pasta, nil
}

func (s *PastaService) Update(ctx context.Context, id uuid.UUID, name, description string, price float64) (*model.Pasta, error) {
	pasta, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	pasta.Name = name
	pasta.Description = &description
	if err := s.repo.Update(ctx, pasta); err != nil {
		return nil, err
	}
	return pasta, nil
}

func (s *PastaService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
