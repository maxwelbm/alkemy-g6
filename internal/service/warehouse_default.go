package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func NewWarehouseDefault(repo models.WarehouseRepository) *WarehouseDefault {
	return &WarehouseDefault{repo: repo}
}

type WarehouseDefault struct {
	repo models.WarehouseRepository
}

func (s *WarehouseDefault) GetAllWarehouses() (v map[int]models.Warehouse, err error) {
	v, err = s.repo.GetAll()
	return
}