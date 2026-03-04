package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PastaRepository interface {
	GetAll(ctx context.Context) ([]model.Pasta, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Pasta, error)
	Search(ctx context.Context, name string) ([]model.Pasta, error)
	Create(ctx context.Context, pasta *model.Pasta) error
	Update(ctx context.Context, pasta *model.Pasta) error
	Delete(ctx context.Context, id uuid.UUID) error
	setCategories(tx *gorm.DB, partID uuid.UUID, categories []model.Category) error
}

type pastaRepository struct {
	db *gorm.DB
}

func NewPastaRepository(db *gorm.DB) PastaRepository {
	return &pastaRepository{db: db}
}

func (r *pastaRepository) GetAll(ctx context.Context) ([]model.Pasta, error) {
	var pastas []model.Pasta
	err := r.db.WithContext(ctx).Preload(clause.Associations).Find(&pastas).Error
	return pastas, err
}

func (r *pastaRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Pasta, error) {
	var pasta model.Pasta
	err := r.db.WithContext(ctx).Preload(clause.Associations).First(&pasta, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pasta, nil
}

func (r *pastaRepository) Search(ctx context.Context, name string) ([]model.Pasta, error) {
	var pastas []model.Pasta
	name = strings.ReplaceAll(name, `\`, `\\`)
	name = strings.ReplaceAll(name, `%`, `\%`)
	name = strings.ReplaceAll(name, `_`, `\_`)
	err := r.db.WithContext(ctx).Where("name ILIKE ? ESCAPE '\\'", "%"+name+"%").Find(&pastas).Error
	return pastas, err
}

func (r *pastaRepository) Create(ctx context.Context, pasta *model.Pasta) error {
	return r.db.WithContext(ctx).Create(pasta).Error
}

func (r *pastaRepository) Update(ctx context.Context, pasta *model.Pasta) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Update simple fields
		if err := tx.Clauses(clause.Returning{}).Save(pasta).Error; err != nil {
			return err
		}
		// Replace many2many relations
		return r.setCategories(tx, pasta.ID, pasta.Categories)
	})
}

func (r *pastaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Pasta{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("part not found")
	}
	return nil
}

func (r *pastaRepository) setCategories(tx *gorm.DB, partID uuid.UUID, categories []model.Category) error {
	return tx.Model(&model.Pasta{ID: partID}).Association("Categories").Replace(categories)
}
