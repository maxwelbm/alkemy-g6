package sectionsrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SectionRepository) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	// Check if the section exists
	var exists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sections WHERE id = ?)", id).Scan(&exists)

	if err != nil {
		return updateSection, err
	}

	if !exists {
		err = models.ErrSectionNotFound
		return updateSection, err
	}

	query := `UPDATE sections SET
		section_number = COALESCE(NULLIF(?, ''), section_number),
		current_temperature = COALESCE(NULLIF(?, ''), current_temperature),
		minimum_temperature = COALESCE(NULLIF(?, ''), minimum_temperature),
		current_capacity = COALESCE(NULLIF(?, ''), current_capacity),
		minimum_capacity = COALESCE(NULLIF(?, ''), minimum_capacity),
		maximum_capacity = COALESCE(NULLIF(?, ''), maximum_capacity),
		warehouse_id = COALESCE(NULLIF(?, ''), warehouse_id),
		product_type_id = COALESCE(NULLIF(?, ''), product_type_id)
	WHERE id = ?`

	_, err = r.db.Exec(query, sec.SectionNumber, sec.CurrentTemperature, sec.MinimumTemperature, sec.CurrentCapacity,
		sec.MinimumCapacity, sec.MaximumCapacity, sec.WarehouseID, sec.ProductTypeID, id)

	// Check for errors
	if err != nil {
		return updateSection, err
	}

	err = r.db.QueryRow(`SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, 
		minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?`, id).Scan(
		&updateSection.ID, &updateSection.SectionNumber, &updateSection.CurrentTemperature,
		&updateSection.MinimumTemperature, &updateSection.CurrentCapacity, &updateSection.MinimumCapacity,
		&updateSection.MaximumCapacity, &updateSection.WarehouseID, &updateSection.ProductTypeID)

	if err != nil {
		return updateSection, err
	}

	return updateSection, nil
}
