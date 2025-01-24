package purchaseordersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type PurchaseOrdersRepositoryMock struct {
	mock.Mock
}

func NewPurchaseOrdersRepositoryMock() *PurchaseOrdersRepositoryMock {
	return &PurchaseOrdersRepositoryMock{}
}

func (m *PurchaseOrdersRepositoryMock) Create(purchaseOrdersDTO models.PurchaseOrdersDTO) (po models.PurchaseOrders, err error) {
	args := m.Called(purchaseOrdersDTO)
	return args.Get(0).(models.PurchaseOrders), args.Error(1)
}
