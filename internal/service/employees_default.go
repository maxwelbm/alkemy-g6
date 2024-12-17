package service

import (
	"github.com/maxwelbm/alkemy-g6b/internal/models"
	employees_repository "github.com/maxwelbm/alkemy-g6b/internal/repository/employees"
)

type EmployeesDefault struct {
	rp employees_repository.Employees
}

func NewEmployeesDefault(rp employees_repository.Employees) *EmployeesDefault {
	return &EmployeesDefault{rp: rp}
}

func (e *EmployeesDefault) GetAllEmployees() (employees map[int]models.Employees, err error) {
	employees, err = e.rp.GetAll()
	return
}
