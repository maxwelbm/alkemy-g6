package product_records_controller

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

type ProductRecordsDefault struct {
	sv models.ProductRecordsService
}

func NewProductRecordsController(sv models.ProductRecordsService) *ProductRecordsDefault {
	return &ProductRecordsDefault{sv: sv}
}

type ProductRecordResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type FullProductRecordJSON struct {
	ID             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}
