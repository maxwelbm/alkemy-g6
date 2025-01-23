package employeesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type expected struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name     string
		id       string
		callErr  error
		expected expected
	}{
		{
			name:    "204 - Successfully deleted an employee",
			id:      "1",
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNoContent,
			},
		},
		{
			name:    "400 - Bad Request error when trying to delete an with invalid (non-numeric) ID",
			id:      "a",
			callErr: nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi:",
			},
		},
		{
			name:    "400 - Bad Request error when trying to delete an with invalid (negative) ID",
			id:      "-10",
			callErr: nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "404 - Not found error when attempting to delete a non-existent employee ID",
			id:      "10",
			callErr: models.ErrEmployeeNotFound,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "employee not found",
			},
		},
		{
			name:    "409 - Conflict error when trying to delete an employee with an associated",
			id:      "1",
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1452",
			},
		},
		{
			name:    "500 - Internal Server Error when trying to delete an employee",
			id:      "1",
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

			r := chi.NewRouter()
			r.Delete("/api/v1/employees/{id}", ctl.Delete)
			url := "/api/v1/employees/" + tt.id

			req := httptest.NewRequest(http.MethodDelete, url, nil)
			res := httptest.NewRecorder()

			sv.On("Delete", mock.AnythingOfType("int")).Return(tt.callErr)
			r.ServeHTTP(res, req)

			var decodedRes struct {
				Message string `json:"message,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "Delete", tt.expected.calls)
			require.NoError(t, err)
			require.Equal(t, tt.expected.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.expected.message)
		})
	}
}
