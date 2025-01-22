package productsctl_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	productsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
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
			name: "201 - Successfully created",
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
		/*{
			name: "400 - body invalid",
			productJSON: `
			{
				"product_code": "P00246",
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
				calls:      1,
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "409 - product code duplicate",
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
			name: "409 - seller_id not exists",
			productJSON: `
			{
				"product_code": "P001",
				"description": "Product 1",
				"height": 10.0,
				"length": 15.0,
				"width": 5.0,
				"weight": 1.0,
				"expiration_rate": 0.1,
				"freezing_rate": 0.3,
				"recommended_freezing_temp": -18.0,
				"product_type_id": 1,
				"seller_id": 1000
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
			name: "422 - seller_id not exists",
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
				calls:      1,
				products:   models.Product{},
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute ProductCode cannot be empty",
			},
		},
		{
			name: "500 - internal server error",
			productJSON: `
			{
				"product_code": "P002",
				"description": "Product 2",
				"height": 10.0,
				"length": 15.0,
				"width": 5.0,
				"weight": 1.0,
				"expiration_rate": 0.1,
				"freezing_rate": 0.3,
				"recommended_freezing_temp": -18.0,
				"product_type_id": 1,
				"seller_id": 1
			}`,
			callErr: errors.New("internal error"),
			expected: expected{
				calls:      1,
				products:   models.Product{},
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
			},
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//product := tt.expected.products
			//Arrange
			sv := service.NewProductsServiceMock()
			ctl := controllers.NewProductsController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/products", strings.NewReader(tt.productJSON))
			res := httptest.NewRecorder()
			//Act
			sv.On("Create", mock.AnythingOfType("models.ProductDTO")).Return(tt.expected.products, tt.callErr)
			ctl.Create(res, req)

			var decodedRes struct {
				Message string                      `json:"message,omitempty"`
				Data    productsctl.ProductFullJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			fmt.Println("response ", res)
			fmt.Println("tt ", tt.expected.products)
			fmt.Println("decode ", decodedRes)

			//Assert
			sv.AssertNumberOfCalls(t, "Create", tt.expected.calls)
			require.Equal(t, tt.expected.statusCode, res.Code)
			if tt.expected.statusCode == http.StatusCreated {
				require.Equal(t, tt.expected.products.ID, decodedRes.Data.ID)
				require.Equal(t, tt.expected.products.ProductCode, decodedRes.Data.ProductCode)
				require.Equal(t, tt.expected.products.Description, decodedRes.Data.Description)
				require.Equal(t, tt.expected.products.Height, decodedRes.Data.Height)
				require.Equal(t, tt.expected.products.Length, decodedRes.Data.Length)
				require.Equal(t, tt.expected.products.Width, decodedRes.Data.Width)
				require.Equal(t, tt.expected.products.NetWeight, decodedRes.Data.NetWeight)
				require.Equal(t, tt.expected.products.ExpirationRate, decodedRes.Data.ExpirationRate)
				require.Equal(t, tt.expected.products.FreezingRate, decodedRes.Data.FreezingRate)
				require.Equal(t, tt.expected.products.RecomFreezTemp, decodedRes.Data.RecomFreezTemp)
				require.Equal(t, tt.expected.products.ProductTypeID, decodedRes.Data.ProductTypeID)
				require.Equal(t, tt.expected.products.SellerID, decodedRes.Data.SellerID)
			}
			require.Contains(t, decodedRes.Message, tt.expected.message)

		})
	}

}
