package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	localities_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/localities"
	sellers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	warehouses_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
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

func NewLocalityController(sv models.LocalityService) *localities_controller.LocalityController {
	return localities_controller.NewLocalityController(sv)
}

func NewWarehousesController(service models.WarehouseService) *warehouses_controller.WarehouseDefault {
	return &warehouses_controller.WarehouseDefault{Service: service}
}
