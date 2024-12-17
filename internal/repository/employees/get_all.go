package repository

import "github.com/maxwelbm/alkemy-g6b/internal/models"

func (e *Employees) GetAll() (employees map[int]models.Employees, err error) {
	employees = make(map[int]models.Employees)

	for key, value := range e.db {
		employees[key] = value
	}
	return
}
