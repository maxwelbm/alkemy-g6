package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/employees"

type Employees struct {
	db map[int]models.Employees
}

func NewEmployees(db map[int]models.Employees) *Employees {
	dataBase := make(map[int]models.Employees)
	if db != nil {
		dataBase = db
	}
	return &Employees{db: dataBase}
}
