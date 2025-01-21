package employeesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	employeesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type expected struct {
		calls      int
		statusCode int
		message    string
		employee   models.Employee
	}
	tests := []struct {
		name         string
		employeeJSON string
		callErr      error
		expected     expected
	}{
		{
			name:         "201 - Successfully created an employee",
			employeeJSON: `{"id": 1, "card_number_id":"3253","first_name":"Rick","last_name":"Grimes","warehouse_id":1}`,
			callErr:      nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusCreated,
				employee:   models.Employee{ID: 1, CardNumberID: "3253", FirstName: "Rick", LastName: "Grimes", WarehouseID: 1},
			},
		},
		{
			name:         "400 - Bad Request error when trying to create an employee",
			employeeJSON: `{"id": 1, "card_number_id":3253,"first_name":"Rick","last_name":"Grimes","warehouse_id":1}`,
			callErr:      nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name:         "409 - Conflict error when trying to create an employee with a duplicate entry",
			employeeJSON: `{"id": 1, "card_number_id":"3253","first_name":"Rick","last_name":"Grimes","warehouse_id":1}`,
			callErr:      &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				employee:   models.Employee{},
			},
		},
		{
			name:         "409 - Conflict error when trying to create an employee with an inexistent warehouseID",
			employeeJSON: `{"id": 1, "card_number_id":"3253","first_name":"Rick","last_name":"Grimes","warehouse_id":10}`,
			callErr:      &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1452",
				employee:   models.Employee{},
			},
		},
		{
			name:         "422 - Conflict error when trying to create an employee with empty parameters",
			employeeJSON: `{"id": 1, "card_number_id":"3923","first_name":"","last_name":"Grimes","warehouse_id":10}`,
			callErr:      nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute FirstName cannot be empty",
				employee:   models.Employee{},
			},
		},
		{
			name:         "422 - Conflict error when trying to create an employee with empty parameters",
			employeeJSON: `{"id": 1,"first_name":"Rick","last_name":"Grimes","warehouse_id":10}`,
			callErr:      nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute CardNumberID cannot be empty",
				employee:   models.Employee{},
			},
		},
		{
			name:         "500 - Internal Server Error when trying to retrieve the list of employees",
			employeeJSON: `{"id": 1, "card_number_id":"3253","first_name":"Rick","last_name":"Grimes","warehouse_id":1}`,
			callErr:      errors.New("internal error"),
			expected: expected{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := service.NewEmployeesServiceMock()
			ctl := controllers.NewEmployeesController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", strings.NewReader(tt.employeeJSON))
			res := httptest.NewRecorder()

			sv.On("Create", mock.AnythingOfType("models.EmployeeDTO")).Return(tt.expected.employee, tt.callErr)
			ctl.Create(res, req)

			var decodedRes struct {
				Message string                        `json:"message,omitempty"`
				Data    employeesctl.EmployeeFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "Create", tt.expected.calls)
			require.NoError(t, err)
			require.Equal(t, tt.expected.statusCode, res.Code)
			if tt.expected.statusCode == http.StatusOK {
				require.Equal(t, tt.expected.employee.ID, decodedRes.Data.ID)
				require.Equal(t, tt.expected.employee.CardNumberID, decodedRes.Data.CardNumberID)
				require.Equal(t, tt.expected.employee.FirstName, decodedRes.Data.FirstName)
				require.Equal(t, tt.expected.employee.LastName, decodedRes.Data.LastName)
				require.Equal(t, tt.expected.employee.WarehouseID, decodedRes.Data.WarehouseID)
			}
			require.Contains(t, decodedRes.Message, tt.expected.message)
		})
	}
}
