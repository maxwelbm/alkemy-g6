package sellers_controller

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

type SellersDefault struct {
	sv models.SellersService
}

func NewSellersController(sv models.SellersService) *SellersDefault {
	return &SellersDefault{sv: sv}
}

type SellerResJSON struct {
	Message string `json:"message,omitempty"`
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
