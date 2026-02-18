package repository

import (
	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"gorm.io/gorm"
)

type PastaRepository interface {
	GetAll() ([]model.Pasta, error)
	GetByID(id uuid.UUID) (*model.Pasta, error)
	Search(name string) ([]model.Pasta, error)
	Create(pasta *model.Pasta) error
	Update(pasta *model.Pasta) error
	Delete(id uuid.UUID) error
}

type pastaRepository struct {
	db *gorm.DB
}

func NewPastaRepository(db *gorm.DB) PastaRepository {
	return &pastaRepository{db: db}
}

func (r *pastaRepository) GetAll() ([]model.Pasta, error) {
	var pastas []model.Pasta
	err := r.db.Find(&pastas).Error
	return pastas, err
}

func (r *pastaRepository) GetByID(id uuid.UUID) (*model.Pasta, error) {
	var pasta model.Pasta
	err := r.db.First(&pasta, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pasta, nil
}

func (r *pastaRepository) Search(name string) ([]model.Pasta, error) {
	var pastas []model.Pasta
	err := r.db.Where("name ILIKE ?", "%"+name+"%").Find(&pastas).Error
	return pastas, err
}

func (r *pastaRepository) Create(pasta *model.Pasta) error {
	return r.db.Create(pasta).Error
}

func (r *pastaRepository) Update(pasta *model.Pasta) error {
	return r.db.Save(pasta).Error
}

func (r *pastaRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Pasta{}, "id = ?", id).Error
}
