package purchaseordersctl

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type PurchaseOrdersController struct {
	sv models.PurchaseOrdersService
}

func NewPurchaseOrdersController(sv models.PurchaseOrdersService) *PurchaseOrdersController {
	return &PurchaseOrdersController{sv: sv}
}

type ResPurchaseOrdersJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type PurchaseOrdersJSON struct {
	OrderNumber     *string `json:"order_number"`
	OrderDate       *string `json:"order_date"`
	TrackingCode    *string `json:"tracking_code"`
	BuyerID         *int    `json:"buyer_id"`
	ProductRecordID *int    `json:"product_record_id"`
}

type PurchaseOrdersResJSON struct {
	ID              int    `json:"id"`
	OrderNumber     string `json:"order_number"`
	OrderDate       string `json:"order_date"`
	TrackingCode    string `json:"tracking_code"`
	BuyerID         int    `json:"buyer_id"`
	ProductRecordID int    `json:"product_record_id"`
}
