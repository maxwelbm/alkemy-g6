package warehouses_repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *WarehouseRepository) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	if warehouse.WarehouseCode != nil {
		query := "SELECT EXISTS(SELECT 1 FROM warehouses WHERE `warehouse_code`=?)"
		var exists bool

		err = r.DB.QueryRow(query, *warehouse.WarehouseCode).Scan(&exists)
		if err != nil {
			return
		}

		if exists {
			err = models.ErrWareHouseCodeExist
			return
		}

		w.WarehouseCode = *warehouse.WarehouseCode
	}

	//COALESCE: used to retain the current value of the field if the new value is null or not applicable
	query := `UPDATE warehouses SET 
				address = COALESCE(NULLIF(?, ''), address), 
				telephone = COALESCE(NULLIF(?, ''), telephone),
				warehouse_code = COALESCE(NULLIF(?, ''), warehouse_code),
				minimum_capacity = COALESCE(NULLIF(?, 0), minimum_capacity),
				minimum_temperature = COALESCE(NULLIF(?, 0), minimum_temperature)
			WHERE id= ?`

	_, err = r.DB.Exec(query, warehouse.Address, warehouse.Telephone, warehouse.WarehouseCode, warehouse.MinimumCapacity, warehouse.MinimumTemperature, id)
	if err != nil {
		return
	}

	query = "SELECT `id`,`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature` FROM warehouses WHERE `id`=?"
	err = r.DB.
		QueryRow(query, id).Scan(&w.Id, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		return
	}

	return
}
