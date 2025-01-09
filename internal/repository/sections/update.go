package sections_repository

import (
	"strings"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SectionRepository) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	fields := []string{}
	values := []interface{}{}

	if sec.SectionNumber != nil {
		fields = append(fields, "section_number = ?")
		values = append(values, *sec.SectionNumber)
	}
	if sec.CurrentTemperature != nil {
		fields = append(fields, "current_temperature = ?")
		values = append(values, *sec.CurrentTemperature)
	}
	if sec.MinimumTemperature != nil {
		fields = append(fields, "minimum_temperature = ?")
		values = append(values, *sec.MinimumTemperature)
	}
	if sec.CurrentCapacity != nil {
		fields = append(fields, "current_capacity = ?")
		values = append(values, *sec.CurrentCapacity)
	}
	if sec.MinimumCapacity != nil {
		fields = append(fields, "minimum_capacity = ?")
		values = append(values, *sec.MinimumCapacity)
	}
	if sec.MaximumCapacity != nil {
		fields = append(fields, "maximum_capacity = ?")
		values = append(values, *sec.MaximumCapacity)
	}
	if sec.WarehouseID != nil {
		fields = append(fields, "warehouse_id = ?")
		values = append(values, *sec.WarehouseID)
	}
	if sec.ProductTypeID != nil {
		fields = append(fields, "product_type_id = ?")
		values = append(values, *sec.ProductTypeID)
	}

	if len(fields) == 0 {
		return
	}

	query := "UPDATE sections SET " + strings.Join(fields, ", ") + " WHERE id = ?"
	values = append(values, id)
	res, err := r.DB.Exec(query, values...)
	if err != nil {
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		err = models.ErrorIdNotFound
		return
	}

	updateSection.ID = id
	if sec.SectionNumber != nil {
		updateSection.SectionNumber = *sec.SectionNumber
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

	return
}
