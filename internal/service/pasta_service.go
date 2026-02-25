package service

import (
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

func (s *PastaService) GetAll() ([]model.Pasta, error) {
	return s.repo.GetAll()
}

func (s *PastaService) GetByID(id uuid.UUID) (*model.Pasta, error) {
	return s.repo.GetByID(id)
}

func (s *PastaService) Search(name string) ([]model.Pasta, error) {
	return s.repo.Search(name)
}

func (s *PastaService) Create(name, description string, price float64) (*model.Pasta, error) {
	pasta := &model.Pasta{
		Name:        name,
		Description: &description,
	}
	if err := s.repo.Create(pasta); err != nil {
		return nil, err
	}
	return pasta, nil
}

func (s *PastaService) Update(id uuid.UUID, name, description string, price float64) (*model.Pasta, error) {
	pasta, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	pasta.Name = name
	pasta.Description = &description
	if err := s.repo.Update(pasta); err != nil {
		return nil, err
	}
	return pasta, nil
}

func (s *PastaService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
