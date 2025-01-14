package sectionsrp

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SectionRepository) GetByID(id int) (sec models.Section, err error) {
	query := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?"

	row := r.DB.QueryRow(query, id)

	err = row.Scan(&sec.ID, &sec.SectionNumber, &sec.CurrentTemperature, &sec.MinimumTemperature, &sec.CurrentCapacity, &sec.MinimumCapacity, &sec.MaximumCapacity, &sec.WarehouseID, &sec.ProductTypeID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrSectionNotFound
			return
		}
		return
	}

	return
}
