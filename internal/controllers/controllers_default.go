package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
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

func NewSellersController(SV models.SellersService) *sellers_controller.SellersDefault {
	return &sellers_controller.SellersDefault{
		SV: SV,
	}
}
func NewWarehousesController(service models.WarehouseService) *warehouses_controller.WarehouseDefault {
	return &warehouses_controller.WarehouseDefault{Service: service}
}
