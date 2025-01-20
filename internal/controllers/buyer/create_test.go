package buyersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	buyersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	buyersrp "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		buyer      models.Buyer
	}
	tests := []struct {
		name      string
		buyerJSON string
		callErr   error
		wanted    wanted
	}{
		{
			name:      "201 - When the buyer is created successfully",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "123456789"}`,
			callErr:   nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
				buyer:      models.Buyer{ID: 1, FirstName: "Buyer 1", LastName: "Ferreira", CardNumberID: "123456789"},
			},
		},
		{
			name:      "400 - When passing a body with invalid json",
			buyerJSON: `{"first_name": 1, "last_name": "Ferreira", "card_number_id": 123456789}`,
			callErr:   nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name:      "422 - When passing a body with a valid json with missing parameters",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira"}`,
			callErr:   nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute CardNumberID cannot be nil",
			},
		},
		{
			name:      "422 - When passing a body with a valid json with empty parameters",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": ""}`,
			callErr:   nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute CardNumberID cannot be empty",
			},
		},
		{
			name:      "409 - When the repository raises a DuplicateEntry error",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "123456789"}`,
			callErr:   &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				buyer:      models.Buyer{},
			},
		},
		{
			name:      "500 - When the repository returns an unexpected error",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "123456789"}`,
			callErr:   errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				buyer:      models.Buyer{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := buyersrp.NewBuyersRepositoryMock()
			sv := service.NewBuyersService(rp)
			ctl := controllers.NewBuyersController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", strings.NewReader(tt.buyerJSON))
			res := httptest.NewRecorder()

			// Act
			rp.On("Create", mock.AnythingOfType("models.BuyerDTO")).Return(tt.wanted.buyer, tt.callErr)
			ctl.Create(res, req)

			// Assert
			var decodedRes struct {
				Message string                  `json:"message,omitempty"`
				Data    buyersctl.FullBuyerJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			rp.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.buyer.FirstName, decodedRes.Data.FirstName)
				require.Equal(t, tt.wanted.buyer.LastName, decodedRes.Data.LastName)
				require.Equal(t, tt.wanted.buyer.CardNumberID, decodedRes.Data.CardNumberID)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
