package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
)

type EmployeesDefault struct {
	rp models.EmployeesService
}

func NewEmployeesDefault(rp models.EmployeesService) *EmployeesDefault {
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

func (e *EmployeesDefault) Update(employees models.EmployeesDTO, id int) (newEmployees models.Employees, err error) {
	newEmployees, err = e.rp.Update(employees, id)
	return
}

func (e *EmployeesDefault) Delete(id int) (err error) {
	err = e.rp.Delete(id)
	return
}
