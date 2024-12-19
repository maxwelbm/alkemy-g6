package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
)

func (r *Warehouses) GetAll() (w []models.Warehouse, err error) {
	for _, value := range r.db {
		w = append(w, value)
	}

	return
}