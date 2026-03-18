package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Part3DModelStatus string

const (
	Part3DModelStatusProcessing Part3DModelStatus = "processing"
	Part3DModelStatusReady      Part3DModelStatus = "ready"
	Part3DModelStatusFailed     Part3DModelStatus = "failed"
)

type Part3DModel struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PartID       uuid.UUID `gorm:"type:uuid"`
	RawURL       string
	ProcessedURL *string
	FileName     string
	FileSize     int64
	Status       Part3DModelStatus
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt
}
