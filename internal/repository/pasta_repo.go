package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"gorm.io/gorm"
)

type PastaRepository interface {
	GetAll(ctx context.Context) ([]model.Pasta, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Pasta, error)
	Search(ctx context.Context, name string) ([]model.Pasta, error)
	Create(ctx context.Context, pasta *model.Pasta) error
	Update(ctx context.Context, pasta *model.Pasta) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type pastaRepository struct {
	db *gorm.DB
}

func NewPastaRepository(db *gorm.DB) PastaRepository {
	return &pastaRepository{db: db}
}

func (r *pastaRepository) GetAll(ctx context.Context) ([]model.Pasta, error) {
	var pastas []model.Pasta
	err := r.db.WithContext(ctx).Find(&pastas).Error
	return pastas, err
}

func (r *pastaRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Pasta, error) {
	var pasta model.Pasta
	err := r.db.WithContext(ctx).First(&pasta, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pasta, nil
}

func (r *pastaRepository) Search(ctx context.Context, name string) ([]model.Pasta, error) {
	var pastas []model.Pasta
	err := r.db.WithContext(ctx).Where("name ILIKE ?", "%"+name+"%").Find(&pastas).Error
	return pastas, err
}

func (r *pastaRepository) Create(ctx context.Context, pasta *model.Pasta) error {
	return r.db.WithContext(ctx).Create(pasta).Error
}

func (r *pastaRepository) Update(ctx context.Context, pasta *model.Pasta) error {
	return r.db.WithContext(ctx).Save(pasta).Error
}

func (r *pastaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Pasta{}, "id = ?", id).Error
}
