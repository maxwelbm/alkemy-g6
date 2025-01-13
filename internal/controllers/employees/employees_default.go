package employees_controller

import (

	models "github.com/maxwelbm/alkemy-g6/internal/models"

)

type EmployeesController struct {
	SV models.EmployeesService
}

func NewEmployeesDefault(sv models.EmployeesService) *EmployeesController {
	return &EmployeesController{SV: sv}
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
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type EmployeesReqJSON struct {
	ID           *int    `json:"id"`
	CardNumberID *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	WarehouseID  *int    `json:"warehouse_id"`
}
