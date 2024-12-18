package repository

import (
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
)

var (
	ErrEmployeesRepositoryNotFound = errors.New("Employees not found")
)

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
