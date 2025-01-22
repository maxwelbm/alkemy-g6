package employeesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	employeesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	type expected struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name      string
		employees []models.Employee
		callErr   error
		expected  expected
	}{
		{
			name:    "200 - Successfully retrieve the list of existing employees",
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
			},
			employees: []models.Employee{
				{ID: 1, CardNumberID: "3253", FirstName: "Rick", LastName: "Grimes", WarehouseID: 1},
				{ID: 2, CardNumberID: "3254", FirstName: "Daryl", LastName: "Dixon", WarehouseID: 1},
				{ID: 3, CardNumberID: "3255", FirstName: "Michonne", LastName: "Hawthorne", WarehouseID: 1},
			},
		},
		{
			name:    "200 - Successfully retrieve empty list of employees",
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
			},
			employees: []models.Employee{},
		},
		{
			name:    "500 - Internal Server Error when trying to retrieve the list of employees",
			callErr: errors.New("internal error"),
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

			req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
			res := httptest.NewRecorder()

			sv.On("GetAll").Return(tt.employees, tt.callErr)
			ctl.GetAll(res, req)

			var decodedRes struct {
				Message string                          `json:"message,omitempty"`
				Data    []employeesctl.EmployeeFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "GetAll", tt.expected.calls)
			require.NoError(t, err)
			require.Equal(t, tt.expected.statusCode, res.Code)
			if len(tt.employees) > 0 {
				for i, employee := range tt.employees {
					require.Equal(t, employee.ID, decodedRes.Data[i].ID)
					require.Equal(t, employee.CardNumberID, decodedRes.Data[i].CardNumberID)
					require.Equal(t, employee.FirstName, decodedRes.Data[i].FirstName)
					require.Equal(t, employee.LastName, decodedRes.Data[i].LastName)
					require.Equal(t, employee.WarehouseID, decodedRes.Data[i].WarehouseID)
				}
			}
			require.Contains(t, decodedRes.Message, tt.expected.message)
		})
	}
}
