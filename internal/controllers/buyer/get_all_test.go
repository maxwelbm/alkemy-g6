package buyersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	buyersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	buyersrp "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
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
		buyers  []models.Buyer
		callErr error
		wanted  wanted
	}{
		{
			name: "200 - When buyers are registered in the database",
			buyers: []models.Buyer{
				{ID: 1, FirstName: "Buyer 1", LastName: "Ferreira", CardNumberID: "123456789"},
				{ID: 2, FirstName: "Buyer 2", LastName: "Ferreira", CardNumberID: "987654321"},
				{ID: 3, FirstName: "Buyer 3", LastName: "Ferreira", CardNumberID: "111111111"},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:    "200 - When no buyers are registered in the database",
			buyers:  []models.Buyer{},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:    "500 - When the repository returns an error",
			buyers:  []models.Buyer{},
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

			req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", nil)
			res := httptest.NewRecorder()

			// Act
			rp.On("GetAll").Return(tt.buyers, tt.callErr)
			ctl.GetAll(res, req)

			var decodedRes struct {
				Message string                    `json:"message,omitempty"`
				Data    []buyersctl.FullBuyerJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			rp.AssertNumberOfCalls(t, "GetAll", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			if len(tt.buyers) > 0 {
				for i, buyer := range tt.buyers {
					require.Equal(t, buyer.FirstName, decodedRes.Data[i].FirstName)
					require.Equal(t, buyer.LastName, decodedRes.Data[i].LastName)
					require.Equal(t, buyer.CardNumberID, decodedRes.Data[i].CardNumberID)
				}
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
