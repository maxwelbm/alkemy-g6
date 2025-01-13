package repository

import (
	"database/sql"
	"errors"
)

var (
	ErrEmployeesRepositoryNotFound       = errors.New("Employees not found")
	ErrEmployeesRepositoryDuplicatedCode = errors.New("Card Number ID already exists")
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
