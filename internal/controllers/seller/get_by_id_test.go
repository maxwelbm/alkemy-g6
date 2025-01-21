package sellersctl_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sellersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		seller     models.Seller
	}
	tests := []struct {
		name    string
		id      string
		callErr error
		wanted  struct {
			calls      int
			statusCode int
			message    string
			seller     models.Seller
		}
	}{
		{
			name:    "200 - When the seller is found",
			id:      "1",
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				seller:     models.Seller{ID: 1, CID: "CID123", CompanyName: "Company 1", Address: "123 Street", Telephone: "123456789", LocalityID: 1},
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
			name:    "404 - When the repository raises a NotFound error",
			id:      "999",
			callErr: models.ErrSellerNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "seller not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewSellersServiceMock()
			ctl := controllers.NewSellersController(sv)

			r := chi.NewRouter()
			r.Get("/api/v1/Sellers/{id}", ctl.GetByID)
			url := "/api/v1/Sellers/" + string(tt.id)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("GetByID", mock.AnythingOfType("int")).Return(tt.wanted.seller, tt.callErr)
			r.ServeHTTP(res, req)
			ctl.GetByID(res, req)

			var decodedRes struct {
				Message string                    `json:"message,omitempty"`
				Data    sellersctl.FullSellerJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "GetByID", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Equal(t, decodedRes.Data.ID, tt.wanted.seller.ID)
			require.Equal(t, decodedRes.Data.CID, tt.wanted.seller.CID)
			require.Equal(t, decodedRes.Data.CompanyName, tt.wanted.seller.CompanyName)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
