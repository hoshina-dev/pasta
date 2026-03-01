package model

func (p *CreatePastaInput) ToModel() *Pasta {
	part := &Pasta{
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

func ApplyUpdatePartInput(part *Pasta, input UpdatePastaInput) {
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
