package models

type WarehouseService interface {
	GetAllWarehouses() (w []Warehouse, err error)
}