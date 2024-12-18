package models

type WarehouseService interface {
	GetAll() (w []Warehouse, err error)
	GetById(id int) (w Warehouse, err error)
}