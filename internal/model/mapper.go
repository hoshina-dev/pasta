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
