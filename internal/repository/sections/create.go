package sectionsrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SectionRepository) Create(sec models.SectionDTO) (newSection models.Section, err error) {
	query := "INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.DB.Exec(query, sec.SectionNumber, sec.CurrentTemperature, sec.MinimumTemperature, sec.CurrentCapacity, sec.MinimumCapacity, sec.MaximumCapacity, sec.WarehouseID, sec.ProductTypeID)
	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	query = "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?"
	err = r.DB.
		QueryRow(query, lastInsertId).
		Scan(&newSection.ID, &newSection.SectionNumber, &newSection.CurrentTemperature, &newSection.MinimumTemperature, &newSection.CurrentCapacity, &newSection.MinimumCapacity, &newSection.MaximumCapacity, &newSection.WarehouseID, &newSection.ProductTypeID)
	if err != nil {
		return
	}

	return
}
