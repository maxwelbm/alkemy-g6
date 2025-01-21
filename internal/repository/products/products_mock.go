package productsrp

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type ProductsRepositoryMock struct {
	mock.Mock
}

func NewProductsRepositoryMock() *ProductsRepositoryMock {
	return &ProductsRepositoryMock{}
}

func (m *ProductsRepositoryMock) GetAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *ProductsRepositoryMock) GetByID(id int) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductsRepositoryMock) Create(prod models.ProductDTO) (models.Product, error) {
	args := m.Called(prod)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductsRepositoryMock) Update(id int, prod models.ProductDTO) (models.Product, error) {
	args := m.Called(id, prod)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductsRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ProductsRepositoryMock) GetReportRecords(id int) ([]models.ProductReportRecords, error) {
	args := m.Called(id)
	return args.Get(0).([]models.ProductReportRecords), args.Error(1)
}
