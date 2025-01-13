package models

type ProductRecord struct {
	ID             int
	LastUpdateDate string
	PurchasePrice  float64
	SalePrice      float64
	ProductId      int
}

type ProductRecordDTO struct {
	ID             int     `json:"id,omitempty"`
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}

type ProductRecordsService interface {
	Create(seller ProductRecordDTO) (sellerReturn ProductRecord, err error)
}

type ProductRecordsRepository interface {
	Create(seller ProductRecordDTO) (sellerReturn ProductRecord, err error)
}
