package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type BuyerDefaultMock struct {
	mock.Mock
}

func NewBuyersServiceMock() *BuyerDefaultMock {
	return &BuyerDefaultMock{}
}

func (m *BuyerDefaultMock) GetAll() ([]models.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]models.Buyer), args.Error(1)
}

func (m *BuyerDefaultMock) GetByID(id int) (models.Buyer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *BuyerDefaultMock) Create(buyer models.BuyerDTO) (models.Buyer, error) {
	args := m.Called(buyer)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *BuyerDefaultMock) Update(id int, buyer models.BuyerDTO) (buyerReturn models.Buyer, err error) {
	args := m.Called(id, buyer)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *BuyerDefaultMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *BuyerDefaultMock) ReportPurchaseOrders(id int) ([]models.BuyerPurchaseOrdersReport, error) {
	args := m.Called(id)
	return args.Get(0).([]models.BuyerPurchaseOrdersReport), args.Error(1)
}
