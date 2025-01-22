package productsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProducts_Delete(t *testing.T) {
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
			name:    "204 - Product deleted successfully",
			id:      "1",
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNoContent,
			},
		},
		{
			name: "400 - Invalid ID format",
			id:   "invalid",
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name: "400 - Negative ID not allowed",
			id:   "-1",
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "404 - Product not found",
			id:      "10",
			callErr: models.ErrProductNotFound,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "product not found",
			},
		},
		{
			name:    "409 - Conflict: Cannot delete or update parent row",
			id:      "1",
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotDeleteOrUpdateParentRow},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1451",
			},
		},
		{
			name:    "500 - Internal server error",
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
			// Arrange
			sv := service.NewProductsServiceMock()
			ctl := controllers.NewProductsController(sv)

			r := chi.NewRouter()
			r.Delete("/api/v1/products/{id}", ctl.Delete)
			req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/"+tt.id, nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("Delete", mock.AnythingOfType("int")).Return(tt.callErr)
			r.ServeHTTP(res, req)
			ctl.Delete(res, req)

			var decodedRes struct {
				Message string `json:"message,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "Delete", tt.expected.calls)
			require.NoError(t, err)
			require.Equal(t, tt.expected.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.expected.message)

		})
	}
}
