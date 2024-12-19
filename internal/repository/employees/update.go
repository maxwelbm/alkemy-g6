package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/employees"

func (e *Employees) Update(employees models.EmployeesDTO, id int) (newEmployees models.Employees, err error) {
	newEmployees, err = e.GetByID(id)
	if err != nil {
		err = ErrEmployeesRepositoryNotFound
		return
	}

	if employees.CardNumberID != "" {
		for _, value := range e.db {
			if value.CardNumberID == employees.CardNumberID && value.ID != employees.ID {
				err = ErrEmployeesRepositoryDuplicatedCode
				return
			}
		}
		newEmployees.CardNumberID = employees.CardNumberID
	}

	if employees.FirstName != "" {
		newEmployees.FirstName = employees.FirstName
	}

	if employees.LastName != "" {
		newEmployees.LastName = employees.LastName
	}

	if employees.WarehouseID != 0 {
		newEmployees.WarehouseID = employees.WarehouseID
	}

	e.db[id] = newEmployees
	return
}
