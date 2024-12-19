package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/employees"

func (e *Employees) GetAll() (employees map[int]models.Employees, err error) {
	employees = make(map[int]models.Employees)

	for key, value := range e.db {
		employees[key] = value
	}
	return
}
