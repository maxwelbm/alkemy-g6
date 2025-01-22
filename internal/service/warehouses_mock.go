package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type WarehousesServiceMock struct {
	mock.Mock
}

func NewWarehousesServiceMock() *WarehousesServiceMock {
	return &WarehousesServiceMock{}
}

func (m *WarehousesServiceMock) GetAll() (w []models.Warehouse, err error) {
	args := m.Called()
	return args.Get(0).([]models.Warehouse), args.Error(1)
}

func (m *WarehousesServiceMock) GetByID(id int) (w models.Warehouse, err error) {
	args := m.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehousesServiceMock) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	args := m.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehousesServiceMock) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	args := m.Called(id, warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehousesServiceMock) Delete(id int) (err error) {
	args := m.Called(id)
	return args.Error(1)
}
