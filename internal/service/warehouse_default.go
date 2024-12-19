package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
)

func NewWarehouseDefault(repo repository.RepoDB) *WarehouseDefault {
	return &WarehouseDefault{repo: repo}
}

type WarehouseDefault struct {
	repo repository.RepoDB
}

func (s *WarehouseDefault) GetAll() (w []models.Warehouse, err error) {
	w, err = s.repo.WarehouseDB.GetAll()
	return
}

func (s *WarehouseDefault) GetById(id int) (w models.Warehouse, err error) {
	w, err = s.repo.WarehouseDB.GetById(id)
	return
}

func (s *WarehouseDefault) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	w, err = s.repo.WarehouseDB.Create(warehouse)
	return
}

func (s *WarehouseDefault) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	w, err = s.repo.WarehouseDB.Update(id, warehouse)
	return
}

func (s *WarehouseDefault) Delete(id int) (err error) {
	err = s.repo.WarehouseDB.Delete(id)
	return
}
