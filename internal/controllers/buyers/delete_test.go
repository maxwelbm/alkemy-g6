package buyersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	buyersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/buyers"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name    string
		id      string
		callErr error
		wanted  wanted
	}{
		{
			name:    "204 - When the buyer is deleted successfully",
			id:      "1",
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNoContent,
			},
		},
		{
			name:    "400 - When passing a non numeric id",
			id:      "abc",
			callErr: nil,
			wanted: wanted{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name:    "400 - When passing a negative id",
			id:      "-1",
			callErr: nil,
			wanted: wanted{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "404 - When the repository raises a BuyerNotFound error",
			id:      "999",
			callErr: models.ErrBuyerNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "buyer not found",
			},
		},
		{
			name:    "500 - When the repository returns an error",
			id:      "1",
			callErr: errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewBuyersServiceMock()
			ctl := controllers.NewBuyersController(sv)

			r := chi.NewRouter()
			r.Delete("/api/v1/buyers/{id}", ctl.Delete)
			url := "/api/v1/buyers/" + string(tt.id)
			req := httptest.NewRequest(http.MethodDelete, url, nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("Delete", mock.AnythingOfType("int")).Return(tt.callErr)
			r.ServeHTTP(res, req)
			ctl.Delete(res, req)

			var decodedRes struct {
				Message string                    `json:"message,omitempty"`
				Data    []buyersctl.FullBuyerJSON `json:"data"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "Delete", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
