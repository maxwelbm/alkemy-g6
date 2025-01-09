package sections_repository

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *SectionRepository) GetAll() (sec []models.Section, err error) {
	query := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections"

	rows, err := r.DB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var section models.Section
		if err = rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID); err != nil {
			return
		}
		sec = append(sec, section)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}
