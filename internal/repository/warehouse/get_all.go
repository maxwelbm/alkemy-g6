package warehouse_repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func (r *Warehouses) GetAll() (w map[int]models.Warehouse, err error) {
	w = make(map[int]models.Warehouse)

	for key, value := range r.db {
		w[key] = value
	}

	return
}