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

type PartRepository interface {
	GetAll(ctx context.Context) ([]model.Part, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Part, error)
	Search(ctx context.Context, name string) ([]model.Part, error)
	Create(ctx context.Context, part *model.Part) error
	Update(ctx context.Context, part *model.Part) error
	Delete(ctx context.Context, id uuid.UUID) error
	setCategories(tx *gorm.DB, partID uuid.UUID, categories []model.Category) error
}

type partRepository struct {
	db *gorm.DB
}

func NewPartRepository(db *gorm.DB) PartRepository {
	return &partRepository{db: db}
}

func (r *partRepository) GetAll(ctx context.Context) ([]model.Part, error) {
	var parts []model.Part
	err := r.db.WithContext(ctx).Preload(clause.Associations).Find(&parts).Error
	return parts, err
}

func (r *partRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Part, error) {
	var part model.Part
	err := r.db.WithContext(ctx).Preload(clause.Associations).First(&part, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &part, nil
}

func (r *partRepository) Search(ctx context.Context, name string) ([]model.Part, error) {
	var parts []model.Part
	name = strings.ReplaceAll(name, `\`, `\\`)
	name = strings.ReplaceAll(name, `%`, `\%`)
	name = strings.ReplaceAll(name, `_`, `\_`)
	err := r.db.WithContext(ctx).Preload(clause.Associations).Where("name ILIKE ? ESCAPE '\\'", "%"+name+"%").Find(&parts).Error
	return parts, err
}

func (r *partRepository) Create(ctx context.Context, part *model.Part) error {
	return r.db.WithContext(ctx).Create(part).Error
}

func (r *partRepository) Update(ctx context.Context, part *model.Part) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Update simple fields
		if err := tx.Clauses(clause.Returning{}).Save(part).Error; err != nil {
			return err
		}
		// Replace many2many relations
		return r.setCategories(tx, part.ID, part.Categories)
	})
}

func (r *partRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Part{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("part not found")
	}
	return nil
}

func (r *partRepository) setCategories(tx *gorm.DB, partID uuid.UUID, categories []model.Category) error {
	return tx.Model(&model.Part{ID: partID}).Association("Categories").Replace(categories)
}
