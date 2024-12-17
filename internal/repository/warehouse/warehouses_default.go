package warehouse_repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func NewWarehouse(db map[int]models.Warehouse) *Warehouses {
	defaultDb := make(map[int]models.Warehouse)
	if db != nil {
		defaultDb = db
	}
	return &Warehouses{db: defaultDb, lastId: 1}
}

type Warehouses struct {
	db 		map[int]models.Warehouse
	lastId	int
}