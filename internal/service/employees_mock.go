package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type EmployeeServiceMock struct {
	mock.Mock
}

func NewEmployeesServiceMock() *EmployeeServiceMock {
	return &EmployeeServiceMock{}
}

func (m *EmployeeServiceMock) GetAll() ([]models.Employee, error) {
	args := m.Called()
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) GetByID(id int) (models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) GetReportInboundOrders(id int) ([]models.EmployeeReportInboundDTO, error) {
	args := m.Called(id)
	return args.Get(0).([]models.EmployeeReportInboundDTO), args.Error(1)
}

func (m *EmployeeServiceMock) GetByCardNumberID(cardNumberID string) (models.Employee, error) {
	args := m.Called(cardNumberID)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Create(buyer models.EmployeeDTO) (models.Employee, error) {
	args := m.Called(buyer)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Update(buyer models.EmployeeDTO, id int) (buyerReturn models.Employee, err error) {
	args := m.Called(id, buyer)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeServiceMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
