package warehousesrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *WarehouseRepository) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	query := "INSERT INTO warehouses (`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query, warehouse.Address, warehouse.Telephone, warehouse.WarehouseCode, warehouse.MinimumCapacity, warehouse.MinimumTemperature)
	if err != nil {
		return
	}

	lastInsertID, err := result.LastInsertID()
	if err != nil {
		return
	}

	query = "SELECT `id`,`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature` FROM warehouses WHERE `id`=?"
	err = r.db.
		QueryRow(query, lastInsertID).Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)

	return
}
