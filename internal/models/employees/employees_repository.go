package models

type EmployeesRepository interface {
	GetAll() (employees map[int]Employees, err error)
	GetByID(id int) (employees Employees, err error)
	Create(employees EmployeesDTO) (newEmployees Employees, err error)
}
