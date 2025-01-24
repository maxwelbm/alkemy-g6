package productbatchesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type ProductBatchesRepositoryMock struct {
	mock.Mock
}

func NewProductBatchesRepositoryMock() *ProductBatchesRepositoryMock {
	return &ProductBatchesRepositoryMock{}
}

func (m *ProductBatchesRepositoryMock) Create(prodBatch models.ProductBatchesDTO) (newProdBatches models.ProductBatches, err error) {
	args := m.Called(prodBatch)
	return args.Get(0).(models.ProductBatches), args.Error(1)
}
