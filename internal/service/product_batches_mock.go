package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type ProductBatchesServiceMock struct {
	mock.Mock
}

func NewProductBatchesServiceMock() *ProductBatchesServiceMock {
	return &ProductBatchesServiceMock{}
}

func (m *ProductBatchesServiceMock) Create(prodBatch models.ProductBatchesDTO) (models.ProductBatches, error) {
	args := m.Called(prodBatch)
	return args.Get(0).(models.ProductBatches), args.Error(1)
}
