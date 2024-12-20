package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/sections"

func (r *Sections) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	updateSection, err = r.GetById(id)
	if err != nil {
		return
	}

	if sec.CurrentTemperature != nil {
		updateSection.CurrentTemperature = *sec.CurrentTemperature
	}
	if sec.MinimumTemperature != nil {
		updateSection.MinimumTemperature = *sec.MinimumTemperature
	}
	if sec.CurrentCapacity != nil {
		updateSection.CurrentCapacity = *sec.CurrentCapacity
	}
	if sec.MinimumCapacity != nil {
		updateSection.MinimumCapacity = *sec.MinimumCapacity
	}
	if sec.MaximumCapacity != nil {
		updateSection.MaximumCapacity = *sec.MaximumCapacity
	}
	if sec.WarehouseID != nil {
		updateSection.WarehouseID = *sec.WarehouseID
	}
	if sec.ProductTypeID != nil {
		updateSection.ProductTypeID = *sec.ProductTypeID
	}

	if sec.SectionNumber != nil {
		for _, s := range r.db {
			if s.SectionNumber == *sec.SectionNumber && s.ID != id {
				err = ErrSectionDuplicatedCode
				return
			}
		}
		updateSection.SectionNumber = *sec.SectionNumber
	}
	r.db[id] = updateSection

	return
}
