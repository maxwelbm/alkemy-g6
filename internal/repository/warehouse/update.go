package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func (r *Warehouses) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	w, err = r.GetById(id)
	if err != nil {
		return
	}
	if warehouse.Address != nil {
		w.Address = *warehouse.Address
	}
	if warehouse.Telephone != nil {
		w.Telephone = *warehouse.Telephone
	}
	if warehouse.MinimumCapacity != nil {
		w.MinimumCapacity = *warehouse.MinimumCapacity
	}
	if warehouse.MinimumTemperature != nil {
		w.MinimumTemperature = *warehouse.MinimumTemperature
	}
	if warehouse.WarehouseCode != nil {
		for _, wh := range r.db {
			if wh.WarehouseCode == *warehouse.WarehouseCode && wh.Id != id {
				err = ErrWarehouseRepositoryDuplicatedCode
				return
			}
		}
		w.WarehouseCode = *warehouse.WarehouseCode
	}
	r.db[id] = w

	return
}