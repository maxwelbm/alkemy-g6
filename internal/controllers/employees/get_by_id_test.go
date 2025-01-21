package employeesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	employeesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetById(t *testing.T) {
	type expected struct {
		calls      int
		statusCode int
		message    string
		employee   models.Employee
	}
	tests := []struct {
		name     string
		id       string
		callErr  error
		expected expected
	}{
		{
			name:    "200 - Successfully retrieve employee details by existing ID",
			id:      "1",
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
				employee:   models.Employee{ID: 1, CardNumberID: "3253", FirstName: "Rick", LastName: "Grimes", WarehouseID: 1},
			},
		},
		{
			name:    "404 - Not found error when attempting to retrieve a non-existent employee ID",
			id:      "10",
			callErr: models.ErrEmployeeNotFound,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "employee not found",
			},
		},
		{
			name:    "400 - Bad Request error when retrieving employee with invalid (non-numeric) ID",
			id:      "a",
			callErr: models.ErrEmployeeNotFound,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi:",
			},
		},
		{
			name:    "400 - Bad Request error when retrieving employee with invalid (negative) ID",
			id:      "-10",
			callErr: models.ErrEmployeeNotFound,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "500 - Internal server error when trying to retrieve a employee by ID",
			callErr: errors.New("internal error"),
			id:      "1",
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

			r := chi.NewRouter()
			r.Get("/api/v1/employees/{id}", ctl.GetByID)
			url := "/api/v1/employees/" + tt.id

			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			sv.On("GetByID", mock.AnythingOfType("int")).Return(tt.expected.employee, tt.callErr)
			r.ServeHTTP(res, req)

			var decodedRes struct {
				Message string                         `json:"message,omitempty"`
				Data    employeesctl.EmployeesFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "GetByID", tt.expected.calls)
			require.NoError(t, err)
			require.Equal(t, tt.expected.statusCode, res.Code)
			require.Equal(t, tt.expected.employee.ID, decodedRes.Data.ID)
			require.Equal(t, tt.expected.employee.CardNumberID, decodedRes.Data.CardNumberID)
			require.Equal(t, tt.expected.employee.FirstName, decodedRes.Data.FirstName)
			require.Equal(t, tt.expected.employee.LastName, decodedRes.Data.LastName)
			require.Equal(t, tt.expected.employee.WarehouseID, decodedRes.Data.WarehouseID)
			require.Contains(t, decodedRes.Message, tt.expected.message)
		})
	}
}
