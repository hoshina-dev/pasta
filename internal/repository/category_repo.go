package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository interface {
	Create(ctx context.Context, m *model.Category) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
	Update(ctx context.Context, m *model.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// Create implements [CategoryRepository].
func (r *categoryRepository) Create(ctx context.Context, m *model.Category) error {
	return r.db.WithContext(ctx).Create(m).Error
}

// Delete implements [CategoryRepository].
func (r *categoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Category{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

// GetAll implements [CategoryRepository].
func (r *categoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var items []model.Category
	err := r.db.WithContext(ctx).Find(&items).Error
	return items, err
}

// GetByID implements [CategoryRepository].
func (r *categoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	var m model.Category
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &m, err
}

// Update implements [CategoryRepository].
func (r *categoryRepository) Update(ctx context.Context, m *model.Category) error {
	return r.db.WithContext(ctx).Clauses(clause.Returning{}).Save(m).Error
}
