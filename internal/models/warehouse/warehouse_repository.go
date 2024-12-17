package models

type WarehouseRepository interface {
	GetAll() (w map[int]Warehouse, err error)
}