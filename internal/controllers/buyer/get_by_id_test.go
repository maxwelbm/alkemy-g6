package buyersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	buyersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	buyersrp "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		buyer      models.Buyer
	}
	tests := []struct {
		name    string
		id      string
		callErr error
		wanted  struct {
			calls      int
			statusCode int
			message    string
			buyer      models.Buyer
		}
	}{
		{
			name:    "200 - When buyers the buyer is found",
			id:      "1",
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				buyer:      models.Buyer{ID: 1, FirstName: "Buyer 1", LastName: "Ferreira", CardNumberID: "123456789"},
			},
		},
		{
			name:    "400 - When passing a non numeric id",
			id:      "abc",
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name:    "400 - When passing a negative id",
			id:      "-1",
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "404 - When the respository raises a NotFound error",
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
			rp := buyersrp.NewBuyersRepositoryMock()
			sv := service.NewBuyersService(rp)
			ctl := controllers.NewBuyersController(sv)

			r := chi.NewRouter()
			r.Get("/api/v1/buyers/{id}", ctl.GetByID)
			url := "/api/v1/buyers/" + string(tt.id)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// Act
			rp.On("GetByID", mock.AnythingOfType("int")).Return(tt.wanted.buyer, tt.callErr)
			r.ServeHTTP(res, req)
			ctl.GetByID(res, req)

			var decodedRes struct {
				Message string                  `json:"message,omitempty"`
				Data    buyersctl.FullBuyerJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			rp.AssertNumberOfCalls(t, "GetByID", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Equal(t, decodedRes.Data.FirstName, tt.wanted.buyer.FirstName)
			require.Equal(t, decodedRes.Data.LastName, tt.wanted.buyer.LastName)
			require.Equal(t, decodedRes.Data.CardNumberID, tt.wanted.buyer.CardNumberID)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
