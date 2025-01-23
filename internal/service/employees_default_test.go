package service_test

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	employeesrp "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var employeesFixture = []models.Employee{
	{
		ID:           1,
		CardNumberID: "1234",
		FirstName:    "Rick",
		LastName:     "Grimes",
		WarehouseID:  1,
	},
	{
		ID:           2,
		CardNumberID: "1235",
		FirstName:    "Daryl",
		LastName:     "Dixon",
		WarehouseID:  1,
	},
	{
		ID:           3,
		CardNumberID: "1236",
		FirstName:    "Michonne",
		LastName:     "Hawthorne",
		WarehouseID:  1,
	},
}

func TestEmployeesDefault_GetAll(t *testing.T) {
	tests := []struct {
		name              string
		employees         []models.Employee
		err               error
		expectedEmployees []models.Employee
		expectedErr       error
	}{
		{
			name:              "Successfully retrieve the list of existing employees",
			employees:         employeesFixture,
			err:               nil,
			expectedEmployees: employeesFixture,
			expectedErr:       nil,
		},
		{
			name:              "Error when trying to retrieve the list of employees",
			employees:         []models.Employee{},
			err:               errors.New("internal error"),
			expectedEmployees: []models.Employee{},
			expectedErr:       errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := employeesrp.NewEmployeesRepositoryMock()
			rp.On("GetAll").Return(tt.employees, tt.err)
			sv := service.NewEmployeesService(rp)

			employee, err := sv.GetAll()

			require.Equal(t, tt.expectedEmployees, employee)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestEmployeesDefault_GetByID(t *testing.T) {
	tests := []struct {
		name             string
		employee         models.Employee
		err              error
		expectedEmployee models.Employee
		expectedErr      error
	}{
		{
			name:             "Successfully retrieve employee details by existing ID",
			employee:         employeesFixture[0],
			err:              nil,
			expectedEmployee: employeesFixture[0],
			expectedErr:      nil,
		},
		{
			name:             "Error when retrieving employee with a non-existent ID",
			employee:         models.Employee{},
			err:              models.ErrEmployeeNotFound,
			expectedEmployee: models.Employee{},
			expectedErr:      models.ErrEmployeeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := employeesrp.NewEmployeesRepositoryMock()
			rp.On("GetByID", tt.employee.ID).Return(tt.employee, tt.err)
			sv := service.NewEmployeesService(rp)

			employee, err := sv.GetByID(tt.employee.ID)

			require.Equal(t, tt.expectedEmployee, employee)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestEmployeesDefault_GetReportInboundOrders(t *testing.T) {
	reports := []models.EmployeeReportInbound{
		{
			ID:           1,
			CardNumberID: "1234",
			FirstName:    "Rick",
			LastName:     "Grimes",
			WarehouseID:  1,
			CountReports: 5,
		},
		{
			ID:           2,
			CardNumberID: "1235",
			FirstName:    "Daryl",
			LastName:     "Dixon",
			WarehouseID:  1,
			CountReports: 3,
		},
		{
			ID:           3,
			CardNumberID: "1236",
			FirstName:    "Michonne",
			LastName:     "Hawthorne",
			WarehouseID:  1,
			CountReports: 2,
		},
	}

	tests := []struct {
		name            string
		id              int
		reports         []models.EmployeeReportInbound
		err             error
		expectedReports []models.EmployeeReportInbound
		expectedErr     error
	}{
		{
			name:            "Successfully retrieve inbound orders'reports for all existing employees",
			id:              0,
			reports:         reports,
			err:             nil,
			expectedReports: reports,
			expectedErr:     nil,
		},
		{
			name:            "Successfully retrieve inbound orders' report for an employee id",
			id:              1,
			reports:         []models.EmployeeReportInbound{reports[0]},
			err:             nil,
			expectedReports: []models.EmployeeReportInbound{reports[0]},
			expectedErr:     nil,
		},
		{
			name:            "Error when retrieving employee with a non-existent ID",
			id:              10,
			reports:         []models.EmployeeReportInbound{},
			err:             models.ErrEmployeeNotFound,
			expectedReports: []models.EmployeeReportInbound{},
			expectedErr:     models.ErrEmployeeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := employeesrp.NewEmployeesRepositoryMock()
			rp.On("GetReportInboundOrders", mock.AnythingOfType("int")).Return(tt.reports, tt.err)
			sv := service.NewEmployeesService(rp)

			reports, err := sv.GetReportInboundOrders(tt.id)

			require.Equal(t, tt.expectedReports, reports)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestEmployeesDefault_Create(t *testing.T) {
	tests := []struct {
		name             string
		employee         models.Employee
		err              error
		expectedEmployee models.Employee
		expectedErr      error
	}{
		{
			name:             "Successfully create a new employee",
			employee:         employeesFixture[2],
			err:              nil,
			expectedEmployee: employeesFixture[2],
			expectedErr:      nil,
		},
		{
			name:             "Error when trying to create a new employee",
			employee:         models.Employee{},
			err:              errors.New("internal error"),
			expectedEmployee: models.Employee{},
			expectedErr:      errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employeeDTO := models.EmployeeDTO{
				ID:           &tt.employee.ID,
				CardNumberID: &tt.employee.CardNumberID,
				FirstName:    &tt.employee.FirstName,
				LastName:     &tt.employee.LastName,
				WarehouseID:  &tt.employee.WarehouseID,
			}

			rp := employeesrp.NewEmployeesRepositoryMock()
			rp.On("Create", employeeDTO).Return(tt.employee, tt.err)
			sv := service.NewEmployeesService(rp)

			employee, err := sv.Create(employeeDTO)

			require.Equal(t, tt.expectedEmployee, employee)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestEmployeesDefault_Update(t *testing.T) {
	tests := []struct {
		name             string
		id               int
		employee         models.Employee
		err              error
		expectedEmployee models.Employee
		expectedErr      error
	}{
		{
			name:             "Successfully updated an employee",
			employee:         employeesFixture[2],
			err:              nil,
			expectedEmployee: employeesFixture[2],
			expectedErr:      nil,
		},
		{
			name:             "Error when trying to update an employee",
			employee:         models.Employee{},
			err:              errors.New("internal error"),
			expectedEmployee: models.Employee{},
			expectedErr:      errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employeeDTO := models.EmployeeDTO{
				ID:           &tt.employee.ID,
				CardNumberID: &tt.employee.CardNumberID,
				FirstName:    &tt.employee.FirstName,
				LastName:     &tt.employee.LastName,
				WarehouseID:  &tt.employee.WarehouseID,
			}

			rp := employeesrp.NewEmployeesRepositoryMock()
			rp.On("Update", employeeDTO, tt.id).Return(tt.employee, tt.err)
			sv := service.NewEmployeesService(rp)

			employee, err := sv.Update(employeeDTO, tt.id)

			require.Equal(t, tt.expectedEmployee, employee)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestEmployeesDefault_Delete(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		err         error
		expectedErr error
	}{
		{
			name:        "Successfully delete an employee",
			id:          1,
			err:         nil,
			expectedErr: nil,
		},
		{
			name:        "Error when trying to delete an employee",
			id:          1,
			err:         errors.New("internal error"),
			expectedErr: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := employeesrp.NewEmployeesRepositoryMock()
			rp.On("Delete", tt.id).Return(tt.err)
			sv := service.NewEmployeesService(rp)

			err := sv.Delete(tt.id)

			require.Equal(t, tt.expectedErr, err)
		})
	}
}
