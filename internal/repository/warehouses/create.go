package warehouses_repository

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *WarehouseRepository) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	query := "INSERT INTO frescos_db.warehouses (`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)"

	result, err := r.DB.Exec(query, warehouse.Address, warehouse.Telephone, warehouse.WarehouseCode, warehouse.MinimumCapacity, warehouse.MinimumTemperature)
	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	query = "SELECT `id`,`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature` FROM frescos_db.warehouses WHERE `id`=?"
	err = r.DB.
		QueryRow(query, lastInsertId).Scan(&w.Id, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		return
	}

	return
}
