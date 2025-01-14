package service

import (
	"errors"

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
	employees, err = e.rp.GetAll()
	return
}

func (e *EmployeesDefault) GetByID(id int) (employees models.Employees, err error) {
	employees, err = e.rp.GetByID(id)
	return
}

func (e *EmployeesDefault) GetReportInboundOrdersByID(id int) (employees []models.EmployeesReportInboundDTO, err error) {
	employees, err = e.rp.GetReportInboundOrdersByID(id)
	return
}

func (e *EmployeesDefault) Create(employees models.EmployeesDTO) (newEmployees models.Employees, err error) {
	newEmployees, err = e.rp.Create(employees)
	return
}

func (e *EmployeesDefault) Update(employees models.EmployeesDTO, id int) (newEmployees models.Employees, err error) {
	newEmployees, err = e.rp.Update(employees, id)
	return
}

func (e *EmployeesDefault) Delete(id int) (err error) {
	err = e.rp.Delete(id)
	return
}
