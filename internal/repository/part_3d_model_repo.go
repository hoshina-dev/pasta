package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"gorm.io/gorm"
)

type Part3DModelRepository interface {
	Create(ctx context.Context, m *model.Part3DModel) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Part3DModel, error)
	GetByPartID(ctx context.Context, partID uuid.UUID) ([]*model.Part3DModel, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string, processedURL *string) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type part3DModelRepository struct {
	db *gorm.DB
}

func NewPart3DModelRepository(db *gorm.DB) *part3DModelRepository {
	return &part3DModelRepository{db: db}
}

func (r *part3DModelRepository) Create(ctx context.Context, m *model.Part3DModel) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *part3DModelRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Part3DModel, error) {
	var m model.Part3DModel
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *part3DModelRepository) GetByPartID(ctx context.Context, partID uuid.UUID) ([]*model.Part3DModel, error) {
	var models []*model.Part3DModel
	err := r.db.WithContext(ctx).Where("part_id = ?", partID).Order("created_at ASC").Find(&models).Error
	return models, err
}

func (r *part3DModelRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string, processedURL *string) error {
	return r.db.WithContext(ctx).Model(&model.Part3DModel{}).
		Where("id = ?", id).
		Updates(map[string]any{"status": status, "processed_url": processedURL}).Error
}

func (r *part3DModelRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Part3DModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("part 3d model not found")
	}
	return nil
}
