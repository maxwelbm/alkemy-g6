package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	localities_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/localities"
	products_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	sections_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	sellers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	warehouses_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(SV models.BuyerService) *buyers_controller.BuyersDefault {
	return &buyers_controller.BuyersDefault{
		SV: SV,
	}
}

func NewProductsController(SV models.ProductService) *products_controller.ProductsDefault {
	return &products_controller.ProductsDefault{
		SV: SV,
	}
}

func NewSectionsController(SV models.SectionService) *sections_controller.SectionsController {
	return &sections_controller.SectionsController{
		SV: SV,
	}
}

func NewSellersController(sv models.SellersService) *sellers_controller.SellersDefault {
	return sellers_controller.NewSellersController(sv)
}

func NewLocalityController(sv models.LocalityService) *localities_controller.LocalityController {
	return localities_controller.NewLocalityController(sv)
}

func NewWarehousesController(service models.WarehouseService) *warehouses_controller.WarehouseDefault {
	return &warehouses_controller.WarehouseDefault{Service: service}
}
