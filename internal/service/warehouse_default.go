package service

import (
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

var (
	ErrWarehouseServiceEmployeesAssociated = errors.New("Cannot delete warehouse: employees are still associated. Please remove or reassign associated sections before deleting.")
	ErrWarehouseServiceSectionsAssociated  = errors.New("Cannot delete warehouse: sections are still associated. Please remove or reassign associated sections before deleting.")
)

func NewWarehousesService(repo models.WarehouseRepository) *WarehouseDefault {
	return &WarehouseDefault{repo: repo}
}

type WarehouseDefault struct {
	repo models.WarehouseRepository
}

func (s *WarehouseDefault) GetAll() (w []models.Warehouse, err error) {
	w, err = s.repo.GetAll()
	return
}

func (s *WarehouseDefault) GetByID(id int) (w models.Warehouse, err error) {
	w, err = s.repo.GetByID(id)
	return
}

func (s *WarehouseDefault) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	w, err = s.repo.Create(warehouse)
	return
}

func (s *WarehouseDefault) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	w, err = s.repo.Update(id, warehouse)
	return
}

func (s *WarehouseDefault) Delete(id int) (err error) {
	err = s.repo.Delete(id)
	return
}
