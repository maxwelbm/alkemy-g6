package service

import (
	"errors"
	"fmt"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

var (
	ErrWareHousesServiceNotFound = errors.New("Warehouse not found")
)

type EmployeesDefault struct {
	rp models.EmployeesRepository
}

func NewEmployeesService(rp models.EmployeesRepository) *EmployeesDefault {
	return &EmployeesDefault{rp: rp}
}

func (e *EmployeesDefault) GetAll() (employees []models.Employees, err error) {
	fmt.Print("Chegou no service")
	employees, err = e.GetAll()
	return
}

func (e *EmployeesDefault) GetByID(id int) (employees models.Employees, err error) {
	employees, err = e.rp.GetByID(id)
	return
}

// TODO VERIFICAR SE O WAREHOUSE DEVE SER CONSULTADO PREVIAMENTE OU SE DEVE SER CONSULTADO DURANTE O CREATE
func (e *EmployeesDefault) Create(employees models.EmployeesDTO) (newEmployees models.Employees, err error) {
	newEmployees, err = e.rp.Create(employees)
	return
}

/*
func (e *EmployeesDefault) Update(employees models.EmployeesDTO, id int) (newEmployees models.Employees, err error) {

		if employees.WarehouseID != nil {
			if _, err = e.repo.WarehouseDB.GetById(*employees.WarehouseID); err != nil {
				err = ErrWareHousesServiceNotFound
				return
			}
		}

		newEmployees, err = e.repo.EmployeesDB.Update(employees, id)
		return
	}
*/
func (e *EmployeesDefault) Delete(id int) (err error) {
	err = e.rp.Delete(id)
	return
}
