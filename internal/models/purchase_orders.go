package models

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
