package models

import "errors"

var (
	ErrEmployeesNotFound = errors.New("Seller not found")
)

type Employees struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
}

type EmployeesDTO struct {
	ID           *int
	CardNumberID *string
	FirstName    *string
	LastName     *string
	WarehouseID  *int
}

type EmployeesReportInboundDTO struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
	CountReports int
}

type EmployeesService interface {
	GetAll() (employees []Employees, err error)
	GetByID(id int) (employees Employees, err error)
	GetReportInboundOrdersByID(id int) (employees []EmployeesReportInboundDTO, err error)
	Create(employees EmployeesDTO) (newEmployees Employees, err error)
	Update(employees EmployeesDTO, id int) (newEmployees Employees, err error)
	Delete(id int) (err error)
}

type EmployeesRepository interface {
	GetAll() (employees []Employees, err error)
	GetByID(id int) (employees Employees, err error)
	GetReportInboundOrdersByID(id int) (employees []EmployeesReportInboundDTO, err error)
	Create(employees EmployeesDTO) (newEmployees Employees, err error)
	Update(employees EmployeesDTO, id int) (newEmployees Employees, err error)
	Delete(id int) (err error)
}
