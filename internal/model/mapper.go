package model

func (p *CreatePartInput) ToModel() *Part {
	part := &Part{
		Name:             p.Name,
		PartNumber:       p.PartNumber,
		ManufacturerID:   p.ManufacturerID,
		Description:      p.Description,
		Condition:        p.Condition,
		TemperatureStage: p.TemperatureStage,
		UserID:           p.UserID,
		OrganizationID:   p.OrganizationID,
		Images:           p.Images,
	}
	if p.IsAvailable != nil {
		part.IsAvailable = *p.IsAvailable
	}
	return part
}

func ApplyUpdatePartInput(part *Part, input UpdatePartInput) {
	if input.Name != nil {
		part.Name = *input.Name
	}
	if input.Description != nil {
		part.Description = input.Description
	}
	if input.Condition != nil {
		part.Condition = *input.Condition
	}
	if input.TemperatureStage != nil {
		part.TemperatureStage = input.TemperatureStage
	}
	if input.IsAvailable != nil {
		part.IsAvailable = *input.IsAvailable
	}
	if input.Images != nil {
		part.Images = input.Images
	}
}
