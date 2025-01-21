package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type ProductsServiceMock struct {
	mock.Mock
}

func NewProductsServiceMock() *ProductsServiceMock {
	return &ProductsServiceMock{}
}

func (m *ProductsServiceMock) GetAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *ProductsServiceMock) GetByID(id int) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductsServiceMock) Create(prod models.ProductDTO) (models.Product, error) {
	args := m.Called(prod)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductsServiceMock) Update(id int, prod models.ProductDTO) (models.Product, error) {
	args := m.Called(id, prod)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductsServiceMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ProductsServiceMock) GetReportRecords(id int) ([]models.ProductReportRecords, error) {
	args := m.Called(id)
	return args.Get(0).([]models.ProductReportRecords), args.Error(1)
}
