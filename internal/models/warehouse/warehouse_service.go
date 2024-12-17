package models

type WarehouseService interface {
	GetAllWarehouses() (w map[int]Warehouse, err error)
}