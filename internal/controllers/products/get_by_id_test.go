package productsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	productsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProducts_GetByID(t *testing.T) {
	type wanted struct {
		calls      int
		products   models.Product
		statusCode int
		message    string
	}
	tests := []struct {
		name    string
		id      string
		callErr error
		wanted  wanted
	}{
		{
			name:    "200 - product found",
			id:      "1",
			callErr: nil,
			wanted: wanted{
				calls: 1,
				products: models.Product{
					ID:          1,
					ProductCode: "P001",
					Description: "Product 1",
					Height:      10.0, Length: 20.0, Width: 30.0,
					NetWeight:      40.0,
					ExpirationRate: 0.1,
					FreezingRate:   0.2,
					RecomFreezTemp: -10.0,
					ProductTypeID:  1,
					SellerID:       1,
				},
				statusCode: http.StatusOK,
			},
		},
		{
			name: "400 - invalid ID format",
			id:   "invalid",
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name: "400 - negative ID",
			id:   "-1",
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "404 - product not found",
			id:      "10",
			callErr: models.ErrProductNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "product not found",
			},
		},
		{
			name:    "500 - internal server error",
			id:      "1",
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
			product := tt.wanted.products
			// Arrange
			sv := service.NewProductsServiceMock()
			ctl := controllers.NewProductsController(sv)

			r := chi.NewRouter()
			r.Get("/api/v1/products/{id}", ctl.GetByID)
			req := httptest.NewRequest(http.MethodGet, "/api/v1/products/"+string(tt.id), nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("GetByID", mock.AnythingOfType("int")).Return(product, tt.callErr)
			r.ServeHTTP(res, req)
			ctl.GetByID(res, req)

			var decodedRes struct {
				Message string                      `json:"message,omitempty"`
				Data    productsctl.ProductFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "GetByID", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Equal(t, decodedRes.Data.ID, product.ID)
			require.Equal(t, decodedRes.Data.ProductCode, product.ProductCode)
			require.Equal(t, decodedRes.Data.Description, product.Description)
			require.Equal(t, decodedRes.Data.Height, product.Height)
			require.Equal(t, decodedRes.Data.Length, product.Length)
			require.Equal(t, decodedRes.Data.Width, product.Width)
			require.Equal(t, decodedRes.Data.NetWeight, product.NetWeight)
			require.Equal(t, decodedRes.Data.ExpirationRate, product.ExpirationRate)
			require.Equal(t, decodedRes.Data.FreezingRate, product.FreezingRate)
			require.Equal(t, decodedRes.Data.RecomFreezTemp, product.RecomFreezTemp)
			require.Equal(t, decodedRes.Data.ProductTypeID, product.ProductTypeID)
			require.Equal(t, decodedRes.Data.SellerID, product.SellerID)

			require.Contains(t, decodedRes.Message, tt.wanted.message)

		})
	}
}
