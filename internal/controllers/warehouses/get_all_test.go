package warehousesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	warehousesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/require"
)

func TestWarehouses_GetAll(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name       string
		warehouses []models.Warehouse
		callErr    error
		wanted     wanted
	}{
		{
			name: "200 - successfully retrieved all warehouses",
			warehouses: []models.Warehouse{
				{ID: 1, Address: "123 Main St", Telephone: "555-1234", WarehouseCode: "WH001", MinimumCapacity: 100, MinimumTemperature: -10},
				{ID: 2, Address: "100 Main Av", Telephone: "555-9876", WarehouseCode: "WH002", MinimumCapacity: 50, MinimumTemperature: -5},
				{ID: 3, Address: "215 Main Ro", Telephone: "555-2830", WarehouseCode: "WH003", MinimumCapacity: 80, MinimumTemperature: 0},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:       "200 - no warehouses found in the database",
			warehouses: []models.Warehouse{},
			callErr:    nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:       "500 - When the repository returns an error",
			warehouses: []models.Warehouse{},
			callErr:    errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := service.NewWarehousesServiceMock()
			ctl := controllers.NewWarehousesController(sv)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
			res := httptest.NewRecorder()

			sv.On("GetAll").Return(tt.warehouses, tt.callErr)
			ctl.GetAll(res, req)

			var decodedRes struct {
				Message string                               `json:"message,omitempty"`
				Data    []warehousesctl.WarehouseDataResJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			require.Equal(t, tt.wanted.statusCode, res.Code)
			if len(tt.warehouses) > 0 {
				for n, warehouse := range tt.warehouses {
					require.Equal(t, warehouse.ID, decodedRes.Data[n].ID)
					require.Equal(t, warehouse.Address, decodedRes.Data[n].Address)
					require.Equal(t, warehouse.Telephone, decodedRes.Data[n].Telephone)
					require.Equal(t, warehouse.WarehouseCode, decodedRes.Data[n].WarehouseCode)
					require.Equal(t, warehouse.MinimumCapacity, decodedRes.Data[n].MinimumCapacity)
					require.Equal(t, warehouse.MinimumTemperature, decodedRes.Data[n].MinimumTemperature)
				}
			}
			if tt.wanted.message != "" {
				require.Contains(t, decodedRes.Message, tt.wanted.message)
			}

		})
	}
}
