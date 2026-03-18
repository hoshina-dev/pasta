package repository

import (
	"context"

	"github.com/hoshina-dev/pasta/internal/model"
	"gorm.io/gorm"
)

type OptimizationJobLogRepository interface {
	Create(ctx context.Context, log *model.OptimizationJobLog) error
	GetByJobID(ctx context.Context, jobID string) (*model.OptimizationJobLog, error)
}

type optimizationJobLogRepository struct {
	db *gorm.DB
}

func NewOptimizationJobLogRepository(db *gorm.DB) *optimizationJobLogRepository {
	return &optimizationJobLogRepository{db: db}
}

func (r *optimizationJobLogRepository) Create(ctx context.Context, log *model.OptimizationJobLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *optimizationJobLogRepository) GetByJobID(ctx context.Context, jobID string) (*model.OptimizationJobLog, error) {
	var log model.OptimizationJobLog
	err := r.db.WithContext(ctx).Where("job_id = ?", jobID).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}
