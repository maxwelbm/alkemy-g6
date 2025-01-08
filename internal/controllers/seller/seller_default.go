package sellers_controller

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

type SellersController struct {
	SV models.SellersService
}

func NewSellerController(SV models.SellersService) *SellersController {
	return &SellersController{SV: SV}
}

type SellerResJSON struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type FullSellerJSON struct {
	ID          int    `json:"id"`
	CID         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}
