package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
)

func (r *Warehouses) GetById(id int) (w models.Warehouse, err error) {
	w, ok := r.db[id]
	if !ok {
		err = ErrWarehouseRepositoryNotFound
	}
	return
}