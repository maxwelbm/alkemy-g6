package models

import "errors"

var (
	ErrInboundOrdersNotFound = errors.New("inbound Orders not found")
)

type InboundOrders struct {
	ID             int
	OrderDate      string
	OrderNumber    int
	EmployeeID     int
	ProductBatchID int
	WarehouseID    int
}

type InboundOrdersDTO struct {
	ID             *int
	OrderDate      *string
	OrderNumber    *int
	EmployeeID     *int
	ProductBatchID *int
	WarehouseID    *int
}

type InboundOrdersService interface {
	Create(inboundOrders InboundOrdersDTO) (newInboundOrders InboundOrders, err error)
}

type InboundOrdersRepository interface {
	Create(inboundOrders InboundOrdersDTO) (newInboundOrders InboundOrders, err error)
}
