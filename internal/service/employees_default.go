package service

import (
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

var (
	ErrWareHousesServiceNotFound = errors.New("warehouse not found")
)

type EmployeesDefault struct {
	rp models.EmployeesRepository
}

func NewEmployeesService(rp models.EmployeesRepository) *EmployeesDefault {
	return &EmployeesDefault{rp: rp}
}

func (e *EmployeesDefault) GetAll() (employees []models.Employee, err error) {
	employees, err = e.rp.GetAll()
	return
}

func (e *EmployeesDefault) GetByID(id int) (employees models.Employee, err error) {
	employees, err = e.rp.GetByID(id)
	return
}

func (e *EmployeesDefault) GetReportInboundOrders(id int) (employees []models.EmployeeReportInboundDTO, err error) {
	employees, err = e.rp.GetReportInboundOrders(id)
	return
}

func (e *EmployeesDefault) Create(employees models.EmployeeDTO) (newEmployees models.Employee, err error) {
	newEmployees, err = e.rp.Create(employees)
	return
}

func (e *EmployeesDefault) Update(employees models.EmployeeDTO, id int) (newEmployees models.Employee, err error) {
	newEmployees, err = e.rp.Update(employees, id)
	return
}

func (e *EmployeesDefault) Delete(id int) (err error) {
	err = e.rp.Delete(id)
	return
}
