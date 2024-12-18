package models

type WarehouseRepository interface {
	GetAll() (w []Warehouse, err error)
}