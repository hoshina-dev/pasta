package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Pasta struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name             string
	PartNumber       string
	ManufacturerID   uuid.UUID
	Manufacturer     Manufacturer
	Description      *string
	Condition        string
	TemperatureStage *string
	IsAvailable      bool
	UserID           uuid.UUID
	OrganizationID   uuid.UUID
	Images           pq.StringArray `gorm:"type:text[];default:'{}'"`
	Categories       []Category     `gorm:"many2many:part_categories;joinForeignKey:PartId"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt
}

func (Pasta) TableName() string {
	return "parts"
}
