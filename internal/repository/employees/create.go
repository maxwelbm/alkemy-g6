package repository

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
)

func (e *Employees) Create(employees models.EmployeesDTO) (newEmployees models.Employees, err error) {
	id := e.lastID + 1

	for _, value := range e.db {
		if employees.CardNumberID == value.CardNumberID {
			err = ErrEmployeesRepositoryDuplicatedCode
			return
		}
	}

	newEmployees = models.Employees{
		ID:           id,
		CardNumberID: employees.CardNumberID,
		FirstName:    employees.FirstName,
		LastName:     employees.LastName,
		WarehouseID:  employees.WarehouseID,
	}
	e.db[id] = newEmployees
	e.lastID = id

	return
}
