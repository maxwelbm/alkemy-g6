package sellersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	sellersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
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
		seller     models.Seller
	}
	tests := []struct {
		name       string
		id         string
		sellerJSON string
		callErr    error
		wanted     wanted
	}{
		{
			name: "200 - When the seller is updated successfully with all fields",
			id:   "1",
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
				statusCode: http.StatusOK,
				message:    "OK",
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
			name: "400 - When passing a non numeric id",
			id:   "abc",
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
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name: "400 - When passing a negative id",
			id:   "-1",
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
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			id:   "1",
			sellerJSON: `{
				"id": 1,
				"cid": 123,
				"company_name": "Meli",
				"address": "123 Main St",
				"telephone": "123-456-7890",
				"locality_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "422 - When passing a body with empty fields",
			id:   "1",
			sellerJSON: `{
				"id": 1,
				"cid": "",
				"company_name": "",
				"address": "",
				"telephone": "",
				"locality_id": 0
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute CID cannot be empty",
			},
		},
		{
			name: "404 - When the repository don't find the seller",
			id:   "1",
			sellerJSON: `{
				"id": 1,
				"cid": "123",
				"company_name": "Meli",
				"address": "123 Main St",
				"telephone": "123-456-7890",
				"locality_id": 1
			}`,
			callErr: models.ErrSellerNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "seller not found",
			},
		},
		{
			name: "500 - When the repository returns an error",
			id:   "1",
			sellerJSON: `{
						"id": 1,
						"cid": "123",
						"company_name": "Meli",
						"address": "123 Main St",
						"telephone": "123-456-7890",
						"locality_id": 1
					}`,
			callErr: models.ErrorNoChangesMade,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusBadRequest,
				message:    "no changes made",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			id:   "1",
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
			},
		},
		{
			name: "500 - When the repository returns an error",
			id:   "1",
			sellerJSON: `{
						"id": 1,
						"cid": "123",
						"company_name": "Meli",
						"address": "123 Main St",
						"telephone": "123-456-7890",
						"locality_id": 1
					}`,
			callErr: errors.New("Generic error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "Generic error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewSellersServiceMock()
			ctl := sellersctl.NewSellersController(sv)

			r := chi.NewRouter()
			r.Patch("/api/v1/sellers/{id}", ctl.Update)
			url := "/api/v1/sellers/" + tt.id
			req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(tt.sellerJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On(
				"Update",
				mock.AnythingOfType("int"),
				mock.AnythingOfType("models.SellerDTO"),
			).Return(tt.wanted.seller, tt.callErr)
			r.ServeHTTP(res, req)

			// Assert
			var decodedRes struct {
				Message string        `json:"message,omitempty"`
				Data    models.Seller `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)
			require.NoError(t, err)

			sv.AssertNumberOfCalls(t, "Update", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.callErr != nil {
				require.Equal(t, tt.wanted.seller.ID, decodedRes.Data.ID)
				require.Equal(t, tt.wanted.seller.CID, decodedRes.Data.CID)
				require.Equal(t, tt.wanted.seller.CompanyName, decodedRes.Data.CompanyName)
			}
		})
	}
}
