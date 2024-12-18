package sellerController

import (
	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
)

type SellerResponse struct {
	Status  int                   `json:"status"`
	Message string                `json:"message,omitempty"`
	Data    []modelsSeller.Seller `json:"data,omitempty"`
}

type SellersResJSON struct {
	Message string              `json:"message"`
	Data    []SellerDataResJSON `json:"data"`
}

type SellerResJSON struct {
	Message string            `json:"message"`
	Data    SellerDataResJSON `json:"data"`
}

type SellerDataResJSON struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerDefault struct {
	sv modelsSeller.SellerService
}

func NewSellerController(sellerService modelsSeller.SellerService) *SellerDefault {
	return &SellerDefault{sv: sellerService}
}
