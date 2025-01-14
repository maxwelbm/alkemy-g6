package buyersctl

import "github.com/maxwelbm/alkemy-g6/internal/models"

type BuyersDefault struct {
	sv models.BuyerService
}

func NewBuyersController(sv models.BuyerService) *BuyersDefault {
	return &BuyersDefault{sv: sv}
}

type BuyerResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type FullBuyerJSON struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
