package sellersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sellersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/models"
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
		seller     models.Seller
	}
	tests := []struct {
		name       string
		sellerJSON string
		callErr    error
		wanted     wanted
	}{
		{
			name: "201 - When the seller is created successfully",
			sellerJSON: `{
				"id": 1,
				"cid": "123",
				"company_name": "Meli",
				"address": "123 Main St",
				"telephone": "123-456-7890",
				"locality_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
				seller: models.Seller{
					ID:          1,
					CID:         "123",
					CompanyName: "Meli",
					Address:     "123 Main St",
					Telephone:   "123-456-7890",
					LocalityID:  1,
				},
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			sellerJSON: `{
				"id": 1,
				"cid": "123",
				"company_name": "Meli",
				"address": "123 Main St",
				"telephone": 1234567890,
				"locality_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
				seller:     models.Seller{},
			},
		},
		{
			name: "422 - When passing a body with a valid json with missing parameters",
			sellerJSON: `{
				"id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: cid is required",
				seller:     models.Seller{},
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			sellerJSON: `{
				"id": 1,
				"cid": "123",
				"company_name": "Meli",
				"address": "123 Main St",
				"telephone": "123-456-7890",
				"locality_id": 1
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				seller:     models.Seller{},
			},
		},
		{
			name: "409 - When the repository raises a CannotAddOrUpdateChildRow error",
			sellerJSON: `{
				"id": 1,
				"cid": "123",
				"company_name": "Meli",
				"address": "123 Main St",
				"telephone": "123-456-7890",
				"locality_id": 1
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1452",
				seller:     models.Seller{},
			},
		},
		{
			name: "500 - When the repository returns an unexpected error",
			sellerJSON: `{
				"id": 1,
				"cid": "123",
				"company_name": "Meli",
				"address": "123 Main St",
				"telephone": "123-456-7890",
				"locality_id": 1
			}`,
			callErr: errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				seller:     models.Seller{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewSellersServiceMock()
			ctl := controllers.NewSellersController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", strings.NewReader(tt.sellerJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On("Create", mock.AnythingOfType("models.SellerDTO")).Return(tt.wanted.seller, tt.callErr)

			sv.On("Create", mock.MatchedBy(func(dto models.SellerDTO) bool {
				if tt.callErr != nil {
					return true
				}
				return (dto.CID == tt.wanted.seller.CID &&
					dto.Address == tt.wanted.seller.Address &&
					dto.CompanyName == tt.wanted.seller.CompanyName &&
					dto.Telephone == tt.wanted.seller.Telephone &&
					dto.LocalityID == tt.wanted.seller.LocalityID)
			})).Return(tt.wanted.seller, tt.callErr)

			ctl.Create(res, req)

			// Assert
			var decodedRes struct {
				Message string                    `json:"message,omitempty"`
				Data    sellersctl.FullSellerJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			sv.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.seller.ID, decodedRes.Data.ID)
				require.Equal(t, tt.wanted.seller.CID, decodedRes.Data.CID)
				require.Equal(t, tt.wanted.seller.CompanyName, decodedRes.Data.CompanyName)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
