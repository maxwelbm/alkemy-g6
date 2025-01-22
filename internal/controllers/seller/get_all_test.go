package sellersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sellersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name    string
		sellers []models.Seller
		callErr error
		wanted  wanted
	}{
		{
			name: "200 - When sellers are registered in the database",
			sellers: []models.Seller{
				{ID: 1, CID: "123", CompanyName: "Acme Corp", Address: "123 Elm St", Telephone: "555-1234", LocalityID: 1},
				{ID: 2, CID: "456", CompanyName: "Globex Inc", Address: "456 Oak St", Telephone: "555-1235", LocalityID: 2},
				{ID: 3, CID: "789", CompanyName: "Soylent Corp", Address: "789 Pine St", Telephone: "555-1236", LocalityID: 3},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:    "200 - When no sellers are registered in the database",
			sellers: []models.Seller{},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:    "500 - When the service returns an error",
			sellers: []models.Seller{},
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
			sv := service.NewSellersServiceMock()
			ctl := controllers.NewSellersController(sv)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers", nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("GetAll").Return(tt.sellers, tt.callErr)
			ctl.GetAll(res, req)

			var decodedRes struct {
				Message string                      `json:"message,omitempty"`
				Data    []sellersctl.FullSellerJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "GetAll", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			if len(tt.sellers) > 0 {
				for i, seller := range tt.sellers {
					require.Equal(t, seller.ID, decodedRes.Data[i].ID)
					require.Equal(t, seller.CID, decodedRes.Data[i].CID)
					require.Equal(t, seller.CompanyName, decodedRes.Data[i].CompanyName)
				}
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
