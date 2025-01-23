package buyersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type BuyerRepositoryMock struct {
	mock.Mock
}

func NewBuyersRepositoryMock() *BuyerRepositoryMock {
	return &BuyerRepositoryMock{}
}

func (m *BuyerRepositoryMock) GetAll() ([]models.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]models.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) GetByID(id int) (models.Buyer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) GetByCardNumberID(CID string) (models.Buyer, error) {
	args := m.Called(CID)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) Create(buyer models.BuyerDTO) (models.Buyer, error) {
	args := m.Called(buyer)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) Update(id int, Buyer models.BuyerDTO) (BuyerReturn models.Buyer, err error) {
	args := m.Called(id, Buyer)
	return args.Get(0).(models.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *BuyerRepositoryMock) ReportPurchaseOrders(id int) ([]models.BuyerPurchaseOrdersReport, error) {
	args := m.Called(id)
	return args.Get(0).([]models.BuyerPurchaseOrdersReport), args.Error(1)
}
