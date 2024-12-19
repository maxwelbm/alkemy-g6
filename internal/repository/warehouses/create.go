package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
)

func (r *Warehouses) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	err = r.validateWarehouseCode(warehouse)
	if err != nil {
		return
	}
	w = models.Warehouse{
		Id: r.LastId + 1,
		Address: *warehouse.Address,
		Telephone: *warehouse.Telephone,
		WarehouseCode: *warehouse.WarehouseCode,
		MinimumCapacity: *warehouse.MinimumCapacity,
		MinimumTemperature: *warehouse.MinimumTemperature,
	}
	r.db[w.Id] = w
	r.LastId = w.Id

	return
}