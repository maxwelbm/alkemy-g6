package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func (r *Warehouses) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	for _, wh := range r.db {
		if wh.WarehouseCode == *warehouse.WarehouseCode {
			err = ErrWarehouseRepositoryDuplicatedCode
			return
		}
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