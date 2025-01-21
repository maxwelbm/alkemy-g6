package productsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	productsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/require"
)

func TestProducts_GetAll(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name     string
		products []models.Product
		callErr  error
		wanted   wanted
	}{
		{
			name: "200 - find all products registered in the database",
			products: []models.Product{
				{ID: 1, ProductCode: "P001", Description: "Product 1", Height: 10.0, Length: 20.0, Width: 30.0, NetWeight: 40.0, ExpirationRate: 0.1, FreezingRate: 0.2, RecomFreezTemp: -10.0, ProductTypeID: 1, SellerID: 1},
				{ID: 2, ProductCode: "P002", Description: "Product 2", Height: 15.0, Length: 25.0, Width: 35.0, NetWeight: 45.0, ExpirationRate: 0.2, FreezingRate: 0.3, RecomFreezTemp: -15.0, ProductTypeID: 2, SellerID: 2},
				{ID: 3, ProductCode: "P003", Description: "Product 3", Height: 20.0, Length: 30.0, Width: 40.0, NetWeight: 50.0, ExpirationRate: 0.3, FreezingRate: 0.4, RecomFreezTemp: -20.0, ProductTypeID: 3, SellerID: 3},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:     "200 - find all when no products registered in the database",
			products: []models.Product{},
			callErr:  nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:     "500 - when the repository returns an error",
			products: []models.Product{},
			callErr:  errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
			},
		},
	}

	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			// Arrange
			sv := service.NewProductsServiceMock()
			ctl := controllers.NewProductsController(sv)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("GetAll").Return(value.products, value.callErr)
			ctl.GetAll(res, req)

			var decodedRes struct {
				Message string                        `json:"message,omitempty"`
				Data    []productsctl.ProductFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			require.NoError(t, err)
			require.Equal(t, value.wanted.statusCode, res.Code)
			if len(value.products) > 0 {
				for i, product := range value.products {
					require.Equal(t, product.ID, decodedRes.Data[i].ID)
					require.Equal(t, product.ProductCode, decodedRes.Data[i].ProductCode)
					require.Equal(t, product.Description, decodedRes.Data[i].Description)
					require.Equal(t, product.Height, decodedRes.Data[i].Height)
					require.Equal(t, product.Length, decodedRes.Data[i].Length)
					require.Equal(t, product.Width, decodedRes.Data[i].Width)
					require.Equal(t, product.NetWeight, decodedRes.Data[i].NetWeight)
					require.Equal(t, product.ExpirationRate, decodedRes.Data[i].ExpirationRate)
					require.Equal(t, product.FreezingRate, decodedRes.Data[i].FreezingRate)
					require.Equal(t, product.RecomFreezTemp, decodedRes.Data[i].RecomFreezTemp)
					require.Equal(t, product.ProductTypeID, decodedRes.Data[i].ProductTypeID)
					require.Equal(t, product.SellerID, decodedRes.Data[i].SellerID)
				}
			}
			if value.wanted.message != "" {
				require.Contains(t, decodedRes.Message, value.wanted.message)
			}
		})
	}
}
