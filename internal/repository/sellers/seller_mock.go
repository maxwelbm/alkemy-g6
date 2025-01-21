package sellersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type SellerRepositoryMock struct {
	mock.Mock
}

func NewSellersRepositoryMock() *SellerRepositoryMock {
	return &SellerRepositoryMock{}
}

func (m *SellerRepositoryMock) GetAll() ([]models.Seller, error) {
	args := m.Called()
	return args.Get(0).([]models.Seller), args.Error(1)
}

func (m *SellerRepositoryMock) GetByID(id int) (models.Seller, error) {
	args := m.Called(id)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *SellerRepositoryMock) GetByCID(CID int) (models.Seller, error) {
	args := m.Called(CID)
	return args.Get(0).(models.Seller), args.Error(1)
}
