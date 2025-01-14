package inboundordersctl

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type InboundOrdersController struct {
	SV models.InboundOrdersService
}

func NewInboundOrdersDefault(sv models.InboundOrdersService) *InboundOrdersController {
	return &InboundOrdersController{SV: sv}
}

type InboundOrdersAttributes struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    int    `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId string `json:"product_batch_id"`
	WarehouseId    int    `json:"warehouse_id"`
}

type InboundOrdersFinalJSON struct {
	Data []InboundOrdersAttributes `json:"data"`
}

type InboundOrdersResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type InboundOrdersReqJSON struct {
	ID             *int    `json:"id"`
	OrderDate      *string `json:"order_date"`
	OrderNumber    *int    `json:"order_number"`
	EmployeeId     *int    `json:"employee_id"`
	ProductBatchId *int    `json:"product_batch_id"`
	WarehouseId    *int    `json:"warehouse_id"`
}
