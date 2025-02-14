package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type InboundOrdersServiceMock struct {
	mock.Mock
}

func NewInboundOrdersServiceMock() *InboundOrdersServiceMock {
	return &InboundOrdersServiceMock{}
}

func (m *InboundOrdersServiceMock) Create(inboundOrders models.InboundOrderDTO) (newInboundOrders models.InboundOrder, err error) {
	args := m.Called(inboundOrders)
	return args.Get(0).(models.InboundOrder), args.Error(1)
}
