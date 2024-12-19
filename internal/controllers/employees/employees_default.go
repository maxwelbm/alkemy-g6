package controller

import (
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

type Employees struct {
	sv service.EmployeesDefault
}

func NewEmployeesDefault(sv service.EmployeesDefault) *Employees {
	return &Employees{sv: sv}
}

type EmployeesAttributes struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

type EmployeesFinalJSON struct {
	Data []EmployeesAttributes `json:"data"`
}

type EmployeesResJSON struct {
	Message string              `json:"message"`
	Data    EmployeesAttributes `json:"data"`
}
