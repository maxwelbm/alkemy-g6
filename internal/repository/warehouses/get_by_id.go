package warehouses_repository

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *WarehouseRepository) GetById(id int) (w models.Warehouse, err error) {
	query := "SELECT `id`, `address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature` FROM frescos_db.warehouses WHERE `id`=?"

	rows := r.DB.QueryRow(query, id)

	if err = rows.Scan(&w.Id, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrWarehouseRepositoryNotFound
			return
		}
	}

	// Check for errors after rows iteration
	if err = rows.Err(); err != nil {
		return
	}

	return
}
