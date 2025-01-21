package employeesctl

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

type EmployeesController struct {
	sv models.EmployeesService
}

func NewEmployeesDefault(sv models.EmployeesService) *EmployeesController {
	return &EmployeesController{sv: sv}
}

type EmployeeFullJSON struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
	CountReports int    `json:"count_reports,omitempty"`
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

type ReportInboundOrdersFullJSON struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
	CountReports int    `json:"count_reports,omitempty"`
}
