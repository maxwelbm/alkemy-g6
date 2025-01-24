package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type ProductRecordServiceMock struct {
	mock.Mock
}

func NewProductRecordsServiceMock() *ProductRecordServiceMock {
	return &ProductRecordServiceMock{}
}

func (m *ProductRecordServiceMock) Create(buyer models.ProductRecordDTO) (models.ProductRecord, error) {
	args := m.Called(buyer)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}
