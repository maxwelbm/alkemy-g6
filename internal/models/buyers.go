package models

import "errors"

var (
	ErrBuyerNotFound = errors.New("Buyer not found ")
)

type Buyer struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
}

type BuyerDTO struct {
	ID           *int    `json:"id,omitempty"`
	CardNumberID *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}
type BuyerPurchaseOrdersReportJSON struct {
	ID                  int    `json:"id,omitempty"`
	CardNumberID        string `json:"card_number_id,omitempty"`
	FirstName           string `json:"first_name,omitempty"`
	LastName            string `json:"last_name,omitempty"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

type BuyerPurchaseOrdersReport struct {
	ID                  int
	CardNumberID        string
	FirstName           string
	LastName            string
	PurchaseOrdersCount int
}

type BuyerResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type BuyerService interface {
	GetAll() (buyers []Buyer, err error)
	GetByID(id int) (buyer Buyer, err error)
	GetByCardNumberID(cardNumberID string) (buyer Buyer, err error)
	Create(buyer BuyerDTO) (buyerReturn Buyer, err error)
	Update(id int, buyer BuyerDTO) (buyerReturn Buyer, err error)
	Delete(id int) (err error)
	ReportPurchaseOrders(id int) (reports []BuyerPurchaseOrdersReport, err error)
}

type BuyerRepository interface {
	GetAll() (buyers []Buyer, err error)
	GetByID(id int) (buyer Buyer, err error)
	GetByCardNumberID(cardNumberID string) (buyer Buyer, err error)
	Create(buyer BuyerDTO) (buyerReturn Buyer, err error)
	Update(id int, buyer BuyerDTO) (buyerReturn Buyer, err error)
	Delete(id int) (err error)
	ReportPurchaseOrders(id int) (reports []BuyerPurchaseOrdersReport, err error)
}
