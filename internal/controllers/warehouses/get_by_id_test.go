package warehousesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	warehousesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWarehouses_GetByID(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		warehouse  models.Warehouse
	}
	tests := []struct {
		name    string
		id      string
		callErr error
		wanted  wanted
	}{
		{
			name:    "200 - warehouse found",
			id:      "1",
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				warehouse: models.Warehouse{
					ID:                 1,
					Address:            "123 Main St",
					Telephone:          "555-1234",
					WarehouseCode:      "WH001",
					MinimumCapacity:    100,
					MinimumTemperature: -10,
				},
			},
		},
		{
			name:    "400 - negative ID",
			id:      "-1",
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "400 - invalid ID format",
			id:      "a",
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name:    "404 - warehouse not found",
			id:      "50",
			callErr: models.ErrWareHouseNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "warehouse not found",
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
			sv := service.NewWarehousesServiceMock()
			ctl := controllers.NewWarehousesController(sv)

			r := chi.NewRouter()
			r.Get("/api/v1/warehouses/{id}", ctl.GetByID)
			req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/"+tt.id, nil)
			res := httptest.NewRecorder()

			sv.On("GetByID", mock.AnythingOfType("int")).Return(tt.wanted.warehouse, tt.callErr)
			r.ServeHTTP(res, req)

			var decodedRes struct {
				Message string                             `json:"message,omitempty"`
				Data    warehousesctl.WarehouseDataResJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "GetByID", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Equal(t, tt.wanted.warehouse.ID, decodedRes.Data.ID)
			require.Equal(t, tt.wanted.warehouse.Address, decodedRes.Data.Address)
			require.Equal(t, tt.wanted.warehouse.Telephone, decodedRes.Data.Telephone)
			require.Equal(t, tt.wanted.warehouse.WarehouseCode, decodedRes.Data.WarehouseCode)
			require.Equal(t, tt.wanted.warehouse.MinimumCapacity, decodedRes.Data.MinimumCapacity)
			require.Equal(t, tt.wanted.warehouse.MinimumTemperature, decodedRes.Data.MinimumTemperature)

			require.Contains(t, decodedRes.Message, tt.wanted.message)

		})
	}
}
