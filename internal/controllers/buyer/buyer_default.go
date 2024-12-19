package buyerController

import (
	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

type BuyerResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message,omitempty"`
	Data    []modelsBuyer.Buyer `json:"data,omitempty"`
}

type BuyersResJSON struct {
	Message string             `json:"message"`
	Data    []BuyerDataResJSON `json:"data"`
}

type BuyerResJSON struct {
	Message string           `json:"message"`
	Data    BuyerDataResJSON `json:"data"`
}

type BuyerDataResJSON struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerDefault struct {
	sv modelsBuyer.BuyerService
}

func NewBuyerController(buyerService modelsBuyer.BuyerService) *BuyerDefault {
	return &BuyerDefault{sv: buyerService}
}
