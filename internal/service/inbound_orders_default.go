package service

import (
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

type InboundOrdersDefault struct {
	rp models.InboundOrderRepository
}

func NewInboundOrdersService(rp models.InboundOrderRepository) *InboundOrdersDefault {
	return &InboundOrdersDefault{rp: rp}
}

func (e *InboundOrdersDefault) Create(inboundOrders models.InboundOrderDTO) (newInboundOrders models.InboundOrder, err error) {
	newInboundOrders, err = e.rp.Create(inboundOrders)
	return
}
