package models

import "errors"

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

type Employee struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
}

type EmployeeDTO struct {
	ID           *int    `json:"id"`
	CardNumberID *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	WarehouseID  *int    `json:"warehouse_id"`
}

type EmployeeReportInboundDTO struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
	CountReports int
}

type EmployeesService interface {
	GetAll() (employees []Employee, err error)
	GetByID(id int) (employees Employee, err error)
	GetReportInboundOrders(id int) (employees []EmployeeReportInboundDTO, err error)
	Create(employees EmployeeDTO) (newEmployees Employee, err error)
	Update(employees EmployeeDTO, id int) (newEmployees Employee, err error)
	Delete(id int) (err error)
}

type EmployeesRepository interface {
	GetAll() (employees []Employee, err error)
	GetByID(id int) (employees Employee, err error)
	GetReportInboundOrders(id int) (employees []EmployeeReportInboundDTO, err error)
	Create(employees EmployeeDTO) (newEmployees Employee, err error)
	Update(employees EmployeeDTO, id int) (newEmployees Employee, err error)
	Delete(id int) (err error)
}
