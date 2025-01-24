package productrecordsrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type ProductRecordsRepositoryMock struct {
	mock.Mock
}

func NewProductRecordsRepositoryMock() *ProductRecordsRepositoryMock {
	return &ProductRecordsRepositoryMock{}
}

func (m *ProductRecordsRepositoryMock) Create(employee models.ProductRecordDTO) (models.ProductRecord, error) {
	args := m.Called(employee)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}
