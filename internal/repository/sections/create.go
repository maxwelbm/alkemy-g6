package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/sections"

func (r *Sections) Create(sec models.SectionDTO) (newSection models.Section, err error) {
	for _, section := range r.db {
		if section.SectionNumber == sec.SectionNumber {
			err = ErrSectionDuplicatedCode
			return
		}
	}

	newSection = models.Section{
		ID:                 r.lastId + 1,
		SectionNumber:      sec.SectionNumber,
		CurrentTemperature: sec.CurrentTemperature,
		MinimumTemperature: sec.MinimumTemperature,
		CurrentCapacity:    sec.CurrentCapacity,
		MinimumCapacity:    sec.MinimumCapacity,
		MaximumCapacity:    sec.MaximumCapacity,
		WarehouseID:        sec.WarehouseID,
		ProductTypeID:      sec.ProductTypeID,
	}

	r.db[newSection.ID] = newSection
	r.lastId = newSection.ID

	return
}
