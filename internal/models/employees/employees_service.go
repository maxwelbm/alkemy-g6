package models

type EmployeesService interface {
	GetAll() (employees map[int]Employees, err error)
}
