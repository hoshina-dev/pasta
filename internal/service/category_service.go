package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/hoshina-dev/pasta/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	c := &model.Category{Name: input.Name}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryService) GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CategoryService) GetAll(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *CategoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
