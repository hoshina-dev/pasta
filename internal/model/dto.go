package model

import "github.com/google/uuid"

type CreatePartInput struct {
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

type UpdatePartInput struct {
	Name             *string
	Description      *string
	Condition        *string
	TemperatureStage *string
	IsAvailable      *bool
	Images           []string    `validate:"omitempty,dive,url"`
	CategoryIDs      []uuid.UUID `validate:"omitempty,dive,uuid4"`
}

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type CreateManufacturerInput struct {
	Name            string  `json:"name"`
	CountryOfOrigin *string `json:"countryOfOrigin,omitempty"`
}
