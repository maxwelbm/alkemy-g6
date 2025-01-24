package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type LocalityDefaultMock struct {
	mock.Mock
}

func NewLocalitiesServiceMock() *LocalityDefaultMock {
	return &LocalityDefaultMock{}
}

func (l *LocalityDefaultMock) Create(locDTO models.LocalityDTO) (loc models.Locality, err error) {
	args := l.Called(locDTO)
	return args.Get(0).(models.Locality), args.Error(1)
}

func (l *LocalityDefaultMock) ReportSellers(id int) (reports []models.LocalitySellersReport, err error) {
	args := l.Called(id)
	return args.Get(0).([]models.LocalitySellersReport), args.Error(1)
}

func (l *LocalityDefaultMock) ReportCarries(id int) (reports []models.LocalityCarriesReport, err error) {
	args := l.Called(id)
	return args.Get(0).([]models.LocalityCarriesReport), args.Error(1)
}
