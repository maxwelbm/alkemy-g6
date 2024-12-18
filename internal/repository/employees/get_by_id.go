package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/employees"

func (r *Employees) GetByID(id int) (employees models.Employees, err error) {
	employees, ok := r.db[id]
	if !ok {
		err = ErrEmployeesRepositoryNotFound
	}
	return
}
