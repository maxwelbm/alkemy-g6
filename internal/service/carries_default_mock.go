package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type CarryDefaultMock struct {
	mock.Mock
}

func NewCarriesServiceMock() *CarryDefaultMock {
	return &CarryDefaultMock{}
}

func (m *CarryDefaultMock) Create(buyer models.CarryDTO) (models.Carry, error) {
	args := m.Called(buyer)
	return args.Get(0).(models.Carry), args.Error(1)
}
