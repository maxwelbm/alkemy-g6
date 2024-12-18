package models

type EmployeesService interface {
	GetAll() (employees map[int]Employees, err error)
	GetByID(id int) (employees Employees, err error)
}
