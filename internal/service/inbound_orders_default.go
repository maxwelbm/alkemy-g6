package service

import (
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

type InboundOrdersDefault struct {
	rp models.InboundOrdersRepository
}

func NewInboundOrdersService(rp models.InboundOrdersRepository) *InboundOrdersDefault {
	return &InboundOrdersDefault{rp: rp}
}

func (e *InboundOrdersDefault) Create(inboundOrders models.InboundOrdersDTO) (newInboundOrders models.InboundOrders, err error) {
	newInboundOrders, err = e.rp.Create(inboundOrders)
	return
}
