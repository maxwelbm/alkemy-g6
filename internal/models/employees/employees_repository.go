package models

type EmployeesRepository interface {
	GetAll() (employees map[int]Employees, err error)
}
