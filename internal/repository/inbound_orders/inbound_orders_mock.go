package inboundordersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type InboundOrdersRepositoryMock struct {
	mock.Mock
}

func NewInboundOrdersRepositoryMock() *InboundOrdersRepositoryMock {
	return &InboundOrdersRepositoryMock{}
}

func (m *InboundOrdersRepositoryMock) Create(inboundOrders models.InboundOrderDTO) (models.InboundOrder, error) {
	args := m.Called(inboundOrders)
	return args.Get(0).(models.InboundOrder), args.Error(1)
}
