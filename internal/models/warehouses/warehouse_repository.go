package models

type WarehouseRepository interface {
	GetAll() (w []Warehouse, err error)
	GetById(id int) (w Warehouse, err error)
	Create(warehouse WarehouseDTO) (w Warehouse, err error)
	Update(id int, warehouse WarehouseDTO) (w Warehouse, err error)
	Delete(id int) (err error)
}