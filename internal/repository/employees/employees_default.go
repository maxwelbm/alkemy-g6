package employeesrp

import (
	"database/sql"
	"errors"
)

var (
	ErrEmployeesRepositoryNotFound       = errors.New("employees not found")
	ErrEmployeesRepositoryDuplicatedCode = errors.New("card Number ID already exists")
)

type EmployeesRepository struct {
	DB *sql.DB
}

func NewEmployeesRepository(DB *sql.DB) *EmployeesRepository {
	repo := &EmployeesRepository{
		DB: DB,
	}
	
	return repo
}
