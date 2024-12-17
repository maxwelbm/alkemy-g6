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

func (s *WarehouseDefault) GetAllWarehouses() (w []models.Warehouse, err error) {
	w, err = s.repo.GetAll()
	return
}