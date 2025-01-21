package employeesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepositoryMock struct {
	mock.Mock
}

func NewEmployeesRepositoryMock() *EmployeeRepositoryMock {
	return &EmployeeRepositoryMock{}
}

func (m *EmployeeRepositoryMock) GetAll() ([]models.Employee, error) {
	args := m.Called()
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) GetByID(id int) (models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) GetReportInboundOrders(id int) ([]models.EmployeeReportInboundDTO, error) {
	args := m.Called(id)
	return args.Get(0).([]models.EmployeeReportInboundDTO), args.Error(1)
}

func (m *EmployeeRepositoryMock) GetByCardNumberID(cardNumberID string) (models.Employee, error) {
	args := m.Called(cardNumberID)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Create(buyer models.EmployeeDTO) (models.Employee, error) {
	args := m.Called(buyer)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Update(buyer models.EmployeeDTO, id int) (buyerReturn models.Employee, err error) {
	args := m.Called(id, buyer)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
