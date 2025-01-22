package productsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	productsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProducts_Create(t *testing.T) {
	type expected struct {
		calls      int
		products   models.Product
		statusCode int
		message    string
	}
	tests := []struct {
		name        string
		productJSON string
		callErr     error
		expected    expected
	}{
		{
			name: "201 - Successfully created product",
			productJSON: `
			{
				"id": 1,
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
				calls: 1,
				products: models.Product{
					ID:             1,
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
				statusCode: http.StatusCreated,
				message:    "Created",
			},
		},
		{
			name: "400 - Invalid request body",
			productJSON: `
			{
				"product_code": 1,
				"description": "X",
				"height": "",
				"length": 15.0,
				"width": 5.0,
				"weight": "1",
				"net_weight": 5,
				"expiration_rate": 0.1,
				"freezing_rate": 0.3,
				"recommended_freezing_temp": -18.0,
				"product_type_id": 1,
				"seller_id": 999
			}`,
			callErr: nil,
			expected: expected{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "409 - Duplicate product code",
			productJSON: `
			{
				"id": 1,
				"product_code": "P011",
				"description": "Product 11",
				"height": 10,
				"length": 20,
				"width": 30,
				"net_weight": 40,
				"expiration_rate": 0.1,
				"freezing_rate": 0.2,
				"recommended_freezing_temp": -10,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			expected: expected{
				calls:      1,
				products:   models.Product{},
				statusCode: http.StatusConflict,
				message:    "1062",
			},
		},
		{
			name: "409 - Seller ID does not exist",
			productJSON: `
			{
				"id": 1,
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
				"seller_id": 1222
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			expected: expected{
				calls:      1,
				products:   models.Product{},
				statusCode: http.StatusConflict,
				message:    "1452",
			},
		},
		{
			name: "422 - Product code cannot be empty",
			productJSON: `
			{
				"product_code": "",
				"description": "Product 1",
				"height": -10.0,
				"length": -15.0,
				"width": -5.0,
				"weight": -1.0,
				"net_weight": 5,
				"expiration_rate": 0.1,
				"freezing_rate": 0.3,
				"recommended_freezing_temp": -18.0,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: nil,
			expected: expected{
				products:   models.Product{},
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute ProductCode cannot be empty",
			},
		},
		{
			name: "422 - Product code cannot be nil",
			productJSON: `
			{
				"product_code": "1",
				"height": 10.0,
				"length": 15.0,
				"width": 5.0,
				"weight": 1.0,
				"net_weight": 5,
				"expiration_rate": 0.1,
				"freezing_rate": 0.3,
				"recommended_freezing_temp": -18.0,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: nil,
			expected: expected{
				products:   models.Product{},
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute Description cannot be nil",
			},
		},
		{
			name: "500 - internal server error",
			productJSON: `
			{
				"id": 1,
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
			product := tt.expected.products
			//Arrange
			sv := service.NewProductsServiceMock()
			ctl := controllers.NewProductsController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/products", strings.NewReader(tt.productJSON))
			res := httptest.NewRecorder()
			//Act
			sv.On("Create", mock.AnythingOfType("models.ProductDTO")).Return(product, tt.callErr)
			ctl.Create(res, req)

			var decodedRes struct {
				Message string                      `json:"message,omitempty"`
				Data    productsctl.ProductFullJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			//Assert
			sv.AssertNumberOfCalls(t, "Create", tt.expected.calls)
			require.Equal(t, tt.expected.statusCode, res.Code)
			if tt.expected.statusCode == http.StatusCreated {
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
