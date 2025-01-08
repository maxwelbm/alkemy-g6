package sellers_controller

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

type SellersDefault struct {
	SV models.SellersService
}

func NewSellersController(SV models.SellersService) *SellersDefault {
	return &SellersDefault{SV: SV}
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
	LocalityID  int    `json:"locality_id"`
}
