package warehousesrp

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *WarehouseRepository) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	if warehouse.WarehouseCode != nil {
		query := "SELECT EXISTS(SELECT 1 FROM warehouses WHERE `warehouse_code`=?)"
		exists := false

		if err = r.db.QueryRow(query, *warehouse.WarehouseCode).Scan(&exists); err != nil {
			return w, err
		}

		if exists {
			err = models.ErrWareHouseCodeExist
			return w, err
		}
	}

	//COALESCE: used to retain the current value of the field if the new value is null or not applicable
	query := `UPDATE warehouses SET 
				address = COALESCE(NULLIF(?, ''), address), 
				telephone = COALESCE(NULLIF(?, ''), telephone),
				warehouse_code = COALESCE(NULLIF(?, ''), warehouse_code),
				minimum_capacity = COALESCE(NULLIF(?, 0), minimum_capacity),
				minimum_temperature = COALESCE(NULLIF(?, 0), minimum_temperature)
			WHERE id= ?`

	_, err = r.db.Exec(query, warehouse.Address, warehouse.Telephone, warehouse.WarehouseCode, warehouse.MinimumCapacity, warehouse.MinimumTemperature, id)
	if err != nil {
		return w, err
	}

	query = "SELECT `id`,`address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature` FROM warehouses WHERE `id`=?"
	if err = r.db.QueryRow(query, id).Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrWareHouseNotFound
			return w, err
		}
		
		return w, err
	}

	return w, err
}
