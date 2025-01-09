package buyers_controller

import "github.com/maxwelbm/alkemy-g6/internal/models"

type BuyersDefault struct {
	SV models.BuyerService
}

func NewBuyersController(SV models.BuyerService) *BuyersDefault {
	return &BuyersDefault{SV: SV}
}

type BuyerResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message,omitempty"`
	Data    []models.Buyer `json:"data,omitempty"`
}

type BuyerResJSON struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type FullBuyerJSON struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
