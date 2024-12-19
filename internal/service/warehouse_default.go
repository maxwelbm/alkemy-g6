package service

import (
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
)

var (
	ErrWarehouseServiceEmployeesAssociated= errors.New("Cannot delete warehouse: employees are still associated. Please remove or reassign associated sections before deleting.")
	ErrWarehouseServiceSectionsAssociated= errors.New("Cannot delete warehouse: sections are still associated. Please remove or reassign associated sections before deleting.")
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
	allSections, err := s.repo.SectionsDB.GetAll()
	if err != nil {
		return
	}
	for _, section := range allSections {
		if section.WarehouseID == id {
			err = ErrWarehouseServiceSectionsAssociated
			return
		}
	}
	allEmployees, err := s.repo.EmployeesDB.GetAll()
	if err != nil {
		return
	}
	for _, employee := range allEmployees {
		if employee.WarehouseID == id {
			err = ErrWarehouseServiceEmployeesAssociated
			return
		}
	}
	err = s.repo.WarehouseDB.Delete(id)
	return
}
