package models

type WarehouseRepository interface {
	GetAll() (w []Warehouse, err error)
	GetById(id int) (w Warehouse, err error)
}