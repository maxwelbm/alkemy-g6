package service

import (
	"github.com/maxwelbm/alkemy-g6b/internal/models"
	repository "github.com/maxwelbm/alkemy-g6b/internal/repository/employees"
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
