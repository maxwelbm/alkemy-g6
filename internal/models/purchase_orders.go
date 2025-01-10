package models

import (
	"errors"
)

var (
	ErrBuyerIDNotExist         = errors.New("Buyer does not exist")
	ErrProductRecordIDNotExist = errors.New("Product Record does not exist")
	ErrOrderNumberExist        = errors.New("Order Number already exist")
)

type PurchaseOrders struct {
	ID              int
	OrderNumber     string
	OrderDate       string
	TrackingCode    string
	BuyerID         int
	ProductRecordID int
}

type PurchaseOrdersDTO struct {
	ID              int
	OrderNumber     string
	OrderDate       string
	TrackingCode    string
	BuyerID         int
	ProductRecordID int
}

type PurchaseOrdersService interface {
	Create(purchaseOrdersDTO PurchaseOrdersDTO) (po PurchaseOrders, err error)
}

type PurchaseOrdersRepository interface {
	Create(purchaseOrdersDTO PurchaseOrdersDTO) (po PurchaseOrders, err error)
}
