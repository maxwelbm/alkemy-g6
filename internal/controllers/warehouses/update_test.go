package warehousesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	warehousesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWarehouses_Update(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		warehouse  models.Warehouse
	}
	tests := []struct {
		name          string
		id            string
		warehouseJSON string
		callErr       error
		wanted        wanted
	}{
		{
			name: "200 - When the warehouse is updated successfully with all fields",
			id:   "1",
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
				statusCode: http.StatusOK,
				message:    "OK",
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
			name: "200 - When the warehouse is updated successfully with missing fields",
			id:   "1",
			warehouseJSON: `{
				"address":"123 Main St",
				"warehouse_code": "WH001", 
				"minimum_capacity":100, 
				"minimum_temperature":-10
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				message:    "OK",
				warehouse: models.Warehouse{
					ID:                 1,
					Address:            "123 Main St",
					WarehouseCode:      "WH001",
					MinimumCapacity:    100,
					MinimumTemperature: -10,
				},
			},
		},
		{
			name: "400 - When passing a negative id",
			id:   "-1",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234",
				"warehouse_code": "WH001",
				"minimum_capacity":100,
				"minimum_temperature":-10
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name: "400 - When passing a non numeric id",
			id:   "a",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234",
				"warehouse_code": "WH001",
				"minimum_capacity":100,
				"minimum_temperature":-10
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			id:   "1",
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
			name: "404 - When the repository raises a WareHouseNotFound error",
			id:   "50",
			warehouseJSON: `{
				"address":"123 Main St",
				"telephone": "555-1234", 
				"warehouse_code": "WH001", 
				"minimum_capacity":100, 
				"minimum_temperature":-10
			}`,
			callErr: models.ErrWareHouseNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "warehouse not found",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			id:   "1",
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
			},
		},
		{
			name: "422 - When passing a body with a valid json but empty parameters",
			id:   "1",
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
			id:   "1",
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
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := service.NewWarehousesServiceMock()
			ctl := controllers.NewWarehousesController(sv)

			r := chi.NewRouter()
			r.Patch("/api/v1/warehouses/{id}", ctl.Update)

			req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/"+tt.id, strings.NewReader(tt.warehouseJSON))
			res := httptest.NewRecorder()

			sv.On(
				"Update",
				mock.AnythingOfType("int"),
				mock.AnythingOfType("models.WarehouseDTO"),
			).Return(tt.wanted.warehouse, tt.callErr)
			r.ServeHTTP(res, req)

			// Assert
			var decodedRes struct {
				Message string                             `json:"message,omitempty"`
				Data    warehousesctl.WarehouseDataResJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			sv.AssertNumberOfCalls(t, "Update", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.callErr != nil {
				require.Equal(t, tt.wanted.warehouse.ID, decodedRes.Data.ID)
				require.Equal(t, tt.wanted.warehouse.Address, decodedRes.Data.Address)
				require.Equal(t, tt.wanted.warehouse.Telephone, decodedRes.Data.Telephone)
				require.Equal(t, tt.wanted.warehouse.WarehouseCode, decodedRes.Data.WarehouseCode)
				require.Equal(t, tt.wanted.warehouse.MinimumCapacity, decodedRes.Data.MinimumCapacity)
				require.Equal(t, tt.wanted.warehouse.MinimumTemperature, decodedRes.Data.MinimumTemperature)
			}
		})
	}
}
