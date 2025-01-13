package models

import "errors"

var (
	ErrorNoChangesMadedInInboundOrders = errors.New("No changes made")
	ErrInboundOrdersNotFound           = errors.New("Inbound Orders not found")
)

type InboundOrders struct {
	ID             int
	OrderDate      string
	OrderNumber    int
	EmployeeId     int
	ProductBatchId string
	WarehouseId    int
}

type InboundOrdersDTO struct {
	ID             *int
	OrderDate      *string
	OrderNumber    *int
	EmployeeId     *int
	ProductBatchId *int
	WarehouseId    *int
}

type InboundOrdersService interface {
	Create(inboundOrders InboundOrdersDTO) (newInboundOrders InboundOrders, err error)
}

type InboundOrdersRepository interface {
	Create(inboundOrders InboundOrdersDTO) (newInboundOrders InboundOrders, err error)
}
