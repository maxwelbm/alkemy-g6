package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type PurchaseOrdersServiceMock struct {
	mock.Mock
}

func NewPurchaseOrdersServiceMock() *PurchaseOrdersServiceMock {
	return &PurchaseOrdersServiceMock{}
}

func (m *PurchaseOrdersServiceMock) Create(purchaseOrdersDTO models.PurchaseOrdersDTO) (po models.PurchaseOrders, err error) {
	args := m.Called(purchaseOrdersDTO)
	return args.Get(0).(models.PurchaseOrders), args.Error(1)
}
