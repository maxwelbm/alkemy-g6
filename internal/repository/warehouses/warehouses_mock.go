package warehousesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type WarehouseRepositoryMock struct {
	mock.Mock
}

func NewWarehouseRepositoryMock() *WarehouseRepositoryMock {
	return &WarehouseRepositoryMock{}
}

func (m *WarehouseRepositoryMock) GetAll() (w []models.Warehouse, err error) {
	args := m.Called()
	return args.Get(0).([]models.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) GetByID(id int) (w models.Warehouse, err error) {
	args := m.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Create(warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	args := m.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Update(id int, warehouse models.WarehouseDTO) (w models.Warehouse, err error) {
	args := m.Called(id, warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Delete(id int) (err error) {
	args := m.Called(id)
	return args.Error(1)
}
