package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func NewWarehouses(db map[int]models.Warehouse) *Warehouses {
	defaultDb := make(map[int]models.Warehouse)
	if db != nil {
		defaultDb = db
	}
	lastId := 0
	for _, w := range db {
		if lastId < w.Id {
			lastId = w.Id
		}
	}
	return &Warehouses{db: defaultDb, lastId: lastId}
}

type Warehouses struct {
	db 		map[int]models.Warehouse
	lastId	int
}