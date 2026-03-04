package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ManufacturerRepository interface {
	Create(ctx context.Context, m *model.Manufacturer) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Manufacturer, error)
	GetAll(ctx context.Context) ([]model.Manufacturer, error)
	Update(ctx context.Context, m *model.Manufacturer) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type manufacturerRepository struct {
	db *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) ManufacturerRepository {
	return &manufacturerRepository{db: db}
}

// Create implements [ManufacturerRepository].
func (r *manufacturerRepository) Create(ctx context.Context, m *model.Manufacturer) error {
	return r.db.WithContext(ctx).Create(m).Error
}

// Delete implements [ManufacturerRepository].
func (r *manufacturerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Manufacturer{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("manufacturer not found")
	}
	return nil
}

// GetAll implements [ManufacturerRepository].
func (r *manufacturerRepository) GetAll(ctx context.Context) ([]model.Manufacturer, error) {
	var items []model.Manufacturer
	err := r.db.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetByID implements [ManufacturerRepository].
func (r *manufacturerRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Manufacturer, error) {
	var m model.Manufacturer
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &m, err
}

// Update implements [ManufacturerRepository].
func (r *manufacturerRepository) Update(ctx context.Context, m *model.Manufacturer) error {
	return r.db.WithContext(ctx).Clauses(clause.Returning{}).Save(m).Error
}
