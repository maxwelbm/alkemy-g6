package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
	"errors"
)

var (
	ErrWarehouseRepositoryNotFound = errors.New("Warehouse not found")
	ErrWarehouseRepositoryDuplicatedCode = errors.New("Warehouse code already exists")
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
	return &Warehouses{db: defaultDb, LastId: lastId}
}

type Warehouses struct {
	db 		map[int]models.Warehouse
	LastId	int
}