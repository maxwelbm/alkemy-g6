package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	warehouses_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(SV models.BuyerService) *buyers_controller.BuyersController {
	return &buyers_controller.BuyersController{
		SV: SV,
	}
}

func NewWarehousesController(service models.WarehouseService) *warehouses_controller.WarehouseDefault {
	return &warehouses_controller.WarehouseDefault{Service: service}
}
