package buyersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	buyersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/buyers"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		buyer      models.Buyer
	}
	tests := []struct {
		name      string
		id        string
		buyerJSON string
		callErr   error
		wanted    wanted
	}{
		{
			name:      "200 - When the buyer is updated successfully with all fields",
			id:        "1",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "123456789"}`,
			callErr:   nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				message:    "OK",
				buyer:      models.Buyer{ID: 1, FirstName: "Buyer 1", LastName: "Ferreira", CardNumberID: "123456789"},
			},
		},
		{
			name:      "200 - When the buyer is updated successfully with missing fields",
			id:        "1",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira"}`,
			callErr:   nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				message:    "OK",
				buyer:      models.Buyer{ID: 1, FirstName: "Buyer 1", LastName: "Ferreira", CardNumberID: "123456789"},
			},
		},
		{
			name:      "400 - When passing a non numeric id",
			id:        "abc",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "123456789"}`,
			callErr:   nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name:      "400 - When passing a negative id",
			id:        "-1",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "123456789"}`,
			callErr:   nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:      "400 - When passing a body with invalid json",
			id:        "1",
			buyerJSON: `{"first_name": 1, "last_name": "Ferreira", "card_number_id": 123456789}`,
			callErr:   nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name:      "404 - When the repository raises a BuyerNotFound error",
			id:        "999",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "12345678"}`,
			callErr:   models.ErrBuyerNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "buyer not found",
			},
		},
		{
			name:      "409 - When the repository raises a DuplicateEntry error",
			id:        "1",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "12345678"}`,
			callErr:   &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
			},
		},
		{
			name:      "422 - When passing a body with empty fields",
			id:        "1",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": ""}`,
			callErr:   nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute CardNumberID cannot be empty",
			},
		},
		{
			name:      "500 - When the repository returns an error",
			id:        "1",
			buyerJSON: `{"first_name": "Buyer 1", "last_name": "Ferreira", "card_number_id": "123456789"}`,
			callErr:   errors.New("internal error"),
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
			ctl := buyersctl.NewBuyersController(sv)

			r := chi.NewRouter()
			r.Patch("/api/v1/buyers/{id}", ctl.Update)
			url := "/api/v1/buyers/" + string(tt.id)
			req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(tt.buyerJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On(
				"Update",
				mock.AnythingOfType("int"),
				mock.AnythingOfType("models.BuyerDTO"),
			).Return(tt.wanted.buyer, tt.callErr)
			r.ServeHTTP(res, req)
			ctl.Update(res, req)

			// Assert
			var decodedRes struct {
				Message string       `json:"message,omitempty"`
				Data    models.Buyer `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)
			require.NoError(t, err)

			sv.AssertNumberOfCalls(t, "Update", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.callErr != nil {
				require.Equal(t, tt.wanted.buyer.FirstName, decodedRes.Data.FirstName)
				require.Equal(t, tt.wanted.buyer.LastName, decodedRes.Data.LastName)
				require.Equal(t, tt.wanted.buyer.CardNumberID, decodedRes.Data.CardNumberID)
			}
		})
	}
}
