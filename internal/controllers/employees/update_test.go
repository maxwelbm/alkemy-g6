package employeesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	employeesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	type expected struct {
		calls      int
		statusCode int
		message    string
		employee   models.Employee
	}
	tests := []struct {
		name         string
		id           string
		employeeJSON string
		callErr      error
		expected     expected
	}{
		{
			name:         "201 - Successfully updated an employee",
			id:           "1",
			employeeJSON: `{"first_name":"Daryl","last_name":"Dixon"}`,
			callErr:      nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
				employee:   models.Employee{ID: 1, CardNumberID: "3253", FirstName: "Daryl", LastName: "Dixon", WarehouseID: 1},
			},
		},
		{
			name:         "400 - Bad Request error when trying to update an employee",
			id:           "1",
			employeeJSON: `{"first_name":1,"last_name":"Dixon"}`,
			callErr:      nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
				employee:   models.Employee{},
			},
		},
		{
			name:         "400 - Bad Request error when trying to update an with invalid (non-numeric) ID",
			id:           "a",
			employeeJSON: `{"first_name":"Daryl,"last_name":"Grimes"}`,
			callErr:      nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi:",
				employee:   models.Employee{},
			},
		},
		{
			name:         "400 - Bad Request error when trying to update an with invalid (negative) ID",
			id:           "-10",
			employeeJSON: `{"first_name":"Daryl,"last_name":"Grimes"}`,
			callErr:      nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
				employee:   models.Employee{},
			},
		},
		{
			name:         "404 - Not found error when attempting to update a non-existent employee ID",
			id:           "10",
			callErr:      models.ErrEmployeeNotFound,
			employeeJSON: `{"first_name":"Daryl","last_name":"Grimes"}`,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "employee not found",
				employee:   models.Employee{},
			},
		},
		{
			name:         "409 - Conflict error when trying to update an employee with a duplicate entry",
			id:           "1",
			employeeJSON: `{"card_number_id":"1"}`,
			callErr:      &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				employee:   models.Employee{},
			},
		},
		{
			name:         "409 - Conflict error when trying to update an employee with an inexistent warehouseID",
			id:           "1",
			employeeJSON: `{"warehouse_id":10}`,
			callErr:      &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1452",
				employee:   models.Employee{},
			},
		},
		{
			name:         "422 - Conflict error when trying to update an employee with empty parameters",
			id:           "1",
			employeeJSON: `{"first_name":"","last_name":"Dixon"}`,
			callErr:      nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute FirstName cannot be empty",
				employee:   models.Employee{},
			},
		},
		{
			name:         "500 - Internal Server Error when trying to update an employee",
			id:           "1",
			employeeJSON: `{"first_name":"Daryl","last_name":"Dixon"}`,
			callErr:      errors.New("internal error"),
			expected: expected{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				employee:   models.Employee{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := service.NewEmployeesServiceMock()
			ctl := controllers.NewEmployeesController(sv)

			r := chi.NewRouter()
			r.Patch("/api/v1/employees/{id}", ctl.Update)
			url := "/api/v1/employees/" + tt.id

			req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(tt.employeeJSON))
			res := httptest.NewRecorder()

			sv.On("Update", mock.AnythingOfType("models.EmployeeDTO"), mock.AnythingOfType("int")).Return(tt.expected.employee, tt.callErr)

			r.ServeHTTP(res, req)

			var decodedRes struct {
				Message string                        `json:"message,omitempty"`
				Data    employeesctl.EmployeeFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "Update", tt.expected.calls)
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
