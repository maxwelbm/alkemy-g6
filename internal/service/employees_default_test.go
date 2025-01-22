package service_test

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	employeesrp "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/require"
)

type Employee struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
}

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
