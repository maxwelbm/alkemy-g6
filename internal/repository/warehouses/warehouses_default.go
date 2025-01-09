package warehouses_repository

import (
	"database/sql"
	"errors"
)

var (
	ErrWarehouseRepositoryNotFound = errors.New("Warehouse not found")
)

type WarehouseRepository struct {
	DB *sql.DB
}

func NewWarehouseRepository(DB *sql.DB) *WarehouseRepository {
	repo := &WarehouseRepository{
		DB: DB,
	}
	return repo
}
