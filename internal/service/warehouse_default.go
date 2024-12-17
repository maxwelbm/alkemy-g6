package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func NewWarehouseDefault(rp models.WarehouseRepository) *WarehouseDefault {
	return &WarehouseDefault{rp: rp}
}

type WarehouseDefault struct {
	rp models.WarehouseRepository
}

func (s *WarehouseDefault) GetAllWarehouses() (v map[int]models.Warehouse, err error) {
	v, err = s.rp.GetAll()
	return
}