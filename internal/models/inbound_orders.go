package models

import "errors"

var (
	ErrInboundOrdersNotFound = errors.New("inbound Orders not found")
)

type InboundOrder struct {
	ID             int
	OrderDate      string
	OrderNumber    int
	EmployeeID     int
	ProductBatchID int
	WarehouseID    int
}

type InboundOrderDTO struct {
	ID             *int
	OrderDate      *string
	OrderNumber    *int
	EmployeeID     *int
	ProductBatchID *int
	WarehouseID    *int
}

type InboundOrderService interface {
	Create(inboundOrders InboundOrderDTO) (newInboundOrders InboundOrder, err error)
}

type InboundOrderRepository interface {
	Create(inboundOrders InboundOrderDTO) (newInboundOrders InboundOrder, err error)
}
