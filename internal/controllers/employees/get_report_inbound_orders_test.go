package employeesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetReportInboundOrders(t *testing.T) {
	type expected struct {
		calls      int
		statusCode int
		message    string
		reports    []models.EmployeeReportInboundDTO
	}
	tests := []struct {
		name     string
		id       string
		callErr  error
		expected expected
	}{
		{
			name:    "200 - Successfully retrieve inbound orders'reports for all existing employees",
			id:      "",
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
				reports: []models.EmployeeReportInboundDTO{
					{ID: 1, CardNumberID: "3253", FirstName: "Rick", LastName: "Grimes", WarehouseID: 1, CountReports: 3},
					{ID: 2, CardNumberID: "3254", FirstName: "Daryl", LastName: "Dixon", WarehouseID: 1, CountReports: 2},
					{ID: 3, CardNumberID: "3255", FirstName: "Carol", LastName: "Peletier", WarehouseID: 1, CountReports: 1},
				},
			},
		},
		{
			name:    "200 - Successfully retrieve inbound orders' report for an employee id",
			id:      "1",
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
				reports: []models.EmployeeReportInboundDTO{
					{ID: 1, CardNumberID: "3253", FirstName: "Rick", LastName: "Grimes", WarehouseID: 1, CountReports: 3},
				},
			},
		},
		{
			name:    "404 - Not found error when attempting to retrieve inbound orders' report of a non-existent employee ID",
			id:      "10",
			callErr: models.ErrEmployeeNotFound,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "employee not found",
				reports:   []models.EmployeeReportInboundDTO{},

			},
		},
		{
			name:    "400 - Bad Request error when retrieving inbound orders' report with invalid (non-numeric) ID",
			id:      "a",
			callErr: models.ErrEmployeeNotFound,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi:",
			},
		},
		{
			name:    "400 - Bad Request error when retrieving inbound orders' report with invalid (negative) ID",
			id:      "-10",
			callErr: models.ErrEmployeeNotFound,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "500 - Internal server error when trying to retrieve inbound orders' report",
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

			var url string
			r := chi.NewRouter()
			if tt.id == "" {
				r.Get("/api/v1/employees/reportInboundOrders", ctl.GetReportInboundOrders)
				url = "/api/v1/employees/reportInboundOrders"
			} else {
				r.Get("/api/v1/employees/reportInboundOrders", ctl.GetReportInboundOrders)
				url = "/api/v1/employees/reportInboundOrders" + "?id=" + tt.id
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			sv.On("GetReportInboundOrders", mock.AnythingOfType("int")).Return(tt.expected.reports, tt.callErr)
			r.ServeHTTP(res, req)

			var decodedRes struct {
				Message string                            `json:"message,omitempty"`
				Data    []models.EmployeeReportInboundDTO `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "GetReportInboundOrders", tt.expected.calls)
			require.NoError(t, err)
			require.Equal(t, tt.expected.statusCode, res.Code)
			for i, report := range decodedRes.Data {
				require.Equal(t, report.ID, decodedRes.Data[i].ID)
				require.Equal(t, report.CardNumberID, decodedRes.Data[i].CardNumberID)
				require.Equal(t, report.FirstName, decodedRes.Data[i].FirstName)
				require.Equal(t, report.LastName, decodedRes.Data[i].LastName)
				require.Equal(t, report.WarehouseID, decodedRes.Data[i].WarehouseID)
				require.Equal(t, report.CountReports, decodedRes.Data[i].CountReports)
			}
			require.Contains(t, decodedRes.Message, tt.expected.message)
		})
	}
}
