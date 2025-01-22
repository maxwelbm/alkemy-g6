package warehousesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	warehousesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWarehouses_Create(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		warehouse  models.Warehouse
	}
	tests := []struct {
		name          string
		warehouseJSON string
		callErr       error
		wanted        wanted
	}{
		{
			name: "201 - When the warehouse is created successfully",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234", 
				"warehouse_code": "WH001", 
				"minimum_capacity":100, 
				"minimum_temperature":-10
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
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
			name: "400 - When passing a body with invalid json",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234",
				"warehouse_code": "WH001",
				"minimum_capacity":"100",
				"minimum_temperature":-10
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234",
				"warehouse_code": "WH001",
				"minimum_capacity":100,
				"minimum_temperature":-10
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				warehouse:  models.Warehouse{},
			},
		},
		{
			name: "422 - When passing a body with a valid json but missing parameters",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234",
				"warehouse_code": "WH001",
				"minimum_temperature":-10
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: the minimum_capacity field cannot be nil",
			},
		},
		{
			name: "422 - When passing a body with a valid json but empty parameters",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234",
				"warehouse_code": "",
				"minimum_capacity":100,
				"minimum_temperature":-10
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: the warehouse_code field cannot be empty",
			},
		},
		{
			name: "500 - When the repository returns an unexpected error",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234",
				"warehouse_code": "WH001",
				"minimum_capacity":100,
				"minimum_temperature":-10
			}`,
			callErr: errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				warehouse:  models.Warehouse{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := service.NewWarehousesServiceMock()
			ctl := controllers.NewWarehousesController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses/", strings.NewReader(tt.warehouseJSON))
			res := httptest.NewRecorder()

			sv.On("Create", mock.AnythingOfType("models.WarehouseDTO")).Return(tt.wanted.warehouse, tt.callErr)
			ctl.Create(res, req)

			// Assert
			var decodedRes struct {
				Message string                             `json:"message,omitempty"`
				Data    warehousesctl.WarehouseDataResJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			sv.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.warehouse.ID, decodedRes.Data.ID)
				require.Equal(t, tt.wanted.warehouse.Address, decodedRes.Data.Address)
				require.Equal(t, tt.wanted.warehouse.Telephone, decodedRes.Data.Telephone)
				require.Equal(t, tt.wanted.warehouse.WarehouseCode, decodedRes.Data.WarehouseCode)
				require.Equal(t, tt.wanted.warehouse.MinimumCapacity, decodedRes.Data.MinimumCapacity)
				require.Equal(t, tt.wanted.warehouse.MinimumTemperature, decodedRes.Data.MinimumTemperature)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)

		})
	}
}
