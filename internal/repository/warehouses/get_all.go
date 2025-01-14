package warehousesrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *WarehouseRepository) GetAll() (w []models.Warehouse, err error) {
	query := "SELECT `id`, `address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature` FROM warehouses"

	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var warehouse models.Warehouse
		if err = rows.Scan(&warehouse.Id, &warehouse.Address, &warehouse.Telephone, &warehouse.WarehouseCode, &warehouse.MinimumCapacity, &warehouse.MinimumTemperature); err != nil {
			return
		}

		w = append(w, warehouse)
	}

	// Check for errors after rows iteration
	if err = rows.Err(); err != nil {
		return
	}

	return
}
