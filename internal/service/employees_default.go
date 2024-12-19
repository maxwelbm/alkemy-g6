package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
)

type EmployeesDefault struct {
	rp repository.Employees
}

func NewEmployeesDefault(rp repository.Employees) *EmployeesDefault {
	return &EmployeesDefault{rp: rp}
}

func (e *EmployeesDefault) GetAll() (employees map[int]models.Employees, err error) {
	employees, err = e.rp.GetAll()
	return
}

func (e *EmployeesDefault) GetByID(id int) (employees models.Employees, err error) {
	employees, err = e.rp.GetByID(id)
	return
}

func (e *EmployeesDefault) Create(employees models.EmployeesDTO) (newEmployees models.Employees, err error) {
	newEmployees, err = e.rp.Create(employees)
	return
}
