package warehouses_repository

import (
	"database/sql"
	"errors"
)

var (
	ErrWarehouseRepositoryNotFound = errors.New("warehouse not found")
)

type WarehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{
		db: db,
	}
}
