package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	sections_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	sellers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(SV models.BuyerService) *buyers_controller.BuyersDefault {
	return &buyers_controller.BuyersDefault{
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
