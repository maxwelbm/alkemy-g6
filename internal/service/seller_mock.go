package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type SellerServiceMock struct {
	mock.Mock
}

func NewSellersServiceMock() *SellerServiceMock {
	return &SellerServiceMock{}
}

func (m *SellerServiceMock) GetAll() ([]models.Seller, error) {
	args := m.Called()
	return args.Get(0).([]models.Seller), args.Error(1)
}

func (m *SellerServiceMock) GetByID(id int) (models.Seller, error) {
	args := m.Called(id)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *SellerServiceMock) GetByCid(CID int) (models.Seller, error) {
	args := m.Called(CID)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *SellerServiceMock) Create(seller models.SellerDTO) (sellerToReturn models.Seller, err error) {
	args := m.Called(seller)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *SellerServiceMock) Update(id int, seller models.SellerDTO) (sellerToReturn models.Seller, err error) {
	args := m.Called(id, seller)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *SellerServiceMock) Delete(id int) (err error) {
	args := m.Called(id)
	return args.Error(1)
}
