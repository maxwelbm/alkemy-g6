package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	products_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(SV models.BuyerService) *buyers_controller.BuyersController {
	return &buyers_controller.BuyersController{
		SV: SV,
	}
}

func NewProductsController(SV models.ProductService) *products_controller.ProductsDefault {
	return &products_controller.ProductsDefault{
		SV: SV,
	}
}
