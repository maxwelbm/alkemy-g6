package models

import "errors"

var (
	ErrWareHouseNotFound  = errors.New("warehouse not found")
	ErrWareHouseCodeExist = errors.New("warehouse code already exist")
)

type Warehouse struct {
	ID                 int     `json:"id"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
}

type WarehouseDTO struct {
	ID                 *int
	Address            *string
	Telephone          *string
	WarehouseCode      *string
	MinimumCapacity    *int
	MinimumTemperature *float64
}

type WarehouseRepository interface {
	GetAll() (w []Warehouse, err error)
	GetByID(id int) (w Warehouse, err error)
	Create(warehouse WarehouseDTO) (w Warehouse, err error)
	Update(id int, warehouse WarehouseDTO) (w Warehouse, err error)
	Delete(id int) (err error)
}

type WarehouseService interface {
	GetAll() (w []Warehouse, err error)
	GetByID(id int) (w Warehouse, err error)
	Create(warehouse WarehouseDTO) (w Warehouse, err error)
	Update(id int, warehouse WarehouseDTO) (w Warehouse, err error)
	Delete(id int) (err error)
}
