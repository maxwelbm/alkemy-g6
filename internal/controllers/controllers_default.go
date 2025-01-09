package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	employees_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(SV models.BuyerService) *buyers_controller.BuyersController {
	return &buyers_controller.BuyersController{
		SV: SV,
	}
}

func NewEmployeesController(SV models.EmployeesService) *employees_controller.EmployeesController {
	return &employees_controller.EmployeesController{
		SV: SV,
	}
}
