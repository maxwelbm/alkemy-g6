package warehousesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type WarehousesRepositoryMock struct {
	mock.Mock
}

func NewWarehousesRepositoryMock() *WarehousesRepositoryMock {
	return &WarehousesRepositoryMock{}
}

func (m *WarehousesRepositoryMock) GetAll() (w []models.Warehouse, err error) {
	args := m.Called()
	return args.Get(0).([]models.Warehouse), args.Error(1)
}

func (m *WarehousesRepositoryMock) GetByID(id int) (w models.Warehouse, err error) {
	args := m.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehousesRepositoryMock) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	args := m.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehousesRepositoryMock) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	args := m.Called(id, warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehousesRepositoryMock) Delete(id int) (err error) {
	args := m.Called(id)
	return args.Error(1)
}
