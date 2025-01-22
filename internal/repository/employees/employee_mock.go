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

func (m *EmployeeRepositoryMock) GetReportInboundOrders(id int) ([]models.EmployeeReportInbound, error) {
	args := m.Called(id)
	return args.Get(0).([]models.EmployeeReportInbound), args.Error(1)
}

func (m *EmployeeRepositoryMock) Create(employee models.EmployeeDTO) (models.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Update(employee models.EmployeeDTO, id int) (employeeReturn models.Employee, err error) {
	args := m.Called(employee, id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
