package model

import "github.com/google/uuid"

type CreatePastaInput struct {
	Name             string    `validate:"required"`
	PartNumber       string    `validate:"required"`
	ManufacturerID   uuid.UUID `validate:"required,uuid4"`
	Description      *string
	Condition        string `validate:"required"`
	TemperatureStage *string
	IsAvailable      *bool
	UserID           uuid.UUID   `validate:"required,uuid4"`
	OrganizationID   uuid.UUID   `validate:"required,uuid4"`
	Images           []string    `validate:"omitempty,dive,url"`
	CategoryIDs      []uuid.UUID `validate:"omitempty,dive,uuid4"`
}

type UpdatePastaInput struct {
	Name             *string
	Description      *string
	Condition        *string
	TemperatureStage *string
	IsAvailable      *bool
	Images           []string    `validate:"omitempty,dive,url"`
	CategoryIDs      []uuid.UUID `validate:"omitempty,dive,uuid4"`
}
