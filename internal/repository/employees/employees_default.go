package repository

import (
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
)

var (
	ErrEmployeesRepositoryNotFound       = errors.New("Employees not found")
	ErrEmployeesRepositoryDuplicatedCode = errors.New("Card Number ID already exists")
)

type Employees struct {
	db     map[int]models.Employees
	lastID int
}

func NewEmployees(db map[int]models.Employees) *Employees {
	dataBase := make(map[int]models.Employees)
	if db != nil {
		dataBase = db
	}

	lastID := 0
	for _, value := range db {
		if lastID < value.ID {
			lastID = value.ID
		}
	}

	return &Employees{db: dataBase, lastID: lastID}
}
