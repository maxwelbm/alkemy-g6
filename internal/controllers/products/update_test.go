package productsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	productsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProducts_Update(t *testing.T) {
	type expected struct {
		calls      int
		statusCode int
		message    string
		product    models.Product
	}
	tests := []struct {
		name        string
		id          string
		productJSON string
		callErr     error
		expected    expected
	}{
		{
			name: "200 - Successfully update product with all fields",
			id:   "1",
			productJSON: `
			{
				"product_code": "P011",
				"description": "Product 11",
				"height": 10,
				"length": 20,
				"width": 30,
				"net_weight": 40,
				"expiration_rate": 1,
				"freezing_rate": 2,
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
				message:    "Updated",
				product: models.Product{
					ProductCode:    "P011",
					Description:    "Product 11",
					Height:         10,
					Length:         20,
					Width:          30,
					NetWeight:      40,
					ExpirationRate: 1,
					FreezingRate:   2,
					RecomFreezTemp: -10,
					ProductTypeID:  1,
					SellerID:       1,
				},
			},
		},
		{
			name: "200 - When updating product with missing fields",
			id:   "1",
			productJSON: `
			{
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: nil,
			expected: expected{
				calls:      1,
				statusCode: http.StatusOK,
				message:    "Updated",
				product: models.Product{
					RecomFreezTemp: -10,
					ProductTypeID:  1,
					SellerID:       1,
				},
			},
		},
		{
			name: "400 - When providing a non-numeric ID",
			id:   "abc",
			productJSON: `
			{
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: nil,
			expected: expected{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name: "400 - When providing a negative ID",
			id:   "-1",
			productJSON: `
			{
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: nil,
			expected: expected{
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name: "400 - When passing a body with invalid JSON",
			id:   "1",
			productJSON: `
			{
				"recommended_freezing_temp": "",
				"product_type_id": "",
				"seller_id": ""
			}`,
			callErr: nil,
			expected: expected{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "404 - When trying to update a product that does not exist",
			id:   "999",
			productJSON: `
			{
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: models.ErrProductNotFound,
			expected: expected{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "product not found",
			},
		},
		{
			name: "409 - When attempting to update to an existing product code",
			id:   "1",
			productJSON: `
			{
				"product_code": "P001",
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: models.ErrProductCodeExist,
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "product code already exist",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			id:   "1",
			productJSON: `
			{
				"product_code": "P001",
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
			},
		},
		{
			name: "409 - When trying to update with an invalid Seller ID",
			id:   "1",
			productJSON: `
			{
				"product_code": "P0221",
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 11111
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			expected: expected{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1452",
			},
		},
		{
			name: "422 - When passing a body with empty fields",
			id:   "1",
			productJSON: `
			{
				"product_code": "P0221",
				"description": "",
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 11111
			}`,
			callErr: nil,
			expected: expected{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute Description cannot be empty",
			},
		},
		{
			name: "500 - When the repository returns an error during update",
			id:   "1",
			productJSON: `
			{
				"product_code": "P0221",
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 11111
			}`,
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
			product := tt.expected.product
			// Arrange
			sv := service.NewProductsServiceMock()
			ctl := productsctl.NewProductsController(sv)

			r := chi.NewRouter()
			r.Patch("/api/v1/products/{id}", ctl.Update)

			req := httptest.NewRequest(http.MethodPatch, "/api/v1/products/"+tt.id, strings.NewReader(tt.productJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("models.ProductDTO")).Return(product, tt.callErr)
			r.ServeHTTP(res, req)

			// Assert
			var decodedRes struct {
				Message string                      `json:"message,omitempty"`
				Data    productsctl.ProductFullJSON `json:"data,omitempty"`
			}

			err := json.NewDecoder(res.Body).Decode(&decodedRes)
			require.NoError(t, err)

			sv.AssertNumberOfCalls(t, "Update", tt.expected.calls)
			require.Equal(t, tt.expected.statusCode, res.Code)
			if tt.callErr != nil {
				require.Equal(t, product.ID, decodedRes.Data.ID)
				require.Equal(t, product.ProductCode, decodedRes.Data.ProductCode)
				require.Equal(t, product.Description, decodedRes.Data.Description)
				require.Equal(t, product.Height, decodedRes.Data.Height)
				require.Equal(t, product.Length, decodedRes.Data.Length)
				require.Equal(t, product.Width, decodedRes.Data.Width)
				require.Equal(t, product.NetWeight, decodedRes.Data.NetWeight)
				require.Equal(t, product.ExpirationRate, decodedRes.Data.ExpirationRate)
				require.Equal(t, product.FreezingRate, decodedRes.Data.FreezingRate)
				require.Equal(t, product.RecomFreezTemp, decodedRes.Data.RecomFreezTemp)
				require.Equal(t, product.ProductTypeID, decodedRes.Data.ProductTypeID)
				require.Equal(t, product.SellerID, decodedRes.Data.SellerID)
			}
			require.Contains(t, decodedRes.Message, tt.expected.message)
		})
	}
}
