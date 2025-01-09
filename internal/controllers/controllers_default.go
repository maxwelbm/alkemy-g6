package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	employees_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	sellers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(SV models.BuyerService) *buyers_controller.BuyersDefault {
	return &buyers_controller.BuyersDefault{
		SV: SV,
	}
}

func NewSellersController(SV models.SellersService) *sellers_controller.SellersDefault {
	return &sellers_controller.SellersDefault{
		SV: SV,
	}
}

func NewEmployeesController(SV models.EmployeesService) *employees_controller.EmployeesController {
	return &employees_controller.EmployeesController{
		SV: SV,
	}
}
