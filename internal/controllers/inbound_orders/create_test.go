package inboundordersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	inboundordersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/inbound_orders"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type wanted struct {
		calls         int
		statusCode    int
		message       string
		inboundOrders models.InboundOrders
	}
	tests := []struct {
		name              string
		inboundOrdersJSON string
		callErr           error
		wanted            wanted
	}{
		{
			name: "201 - When the inbound is created successfully",
			inboundOrdersJSON: `{
				"id": 1,
				"order_date": "2023-10-01",
				"order_number": 12345,
				"employee_id": 1,
				"product_batch_id": 1,
				"warehouse_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
				inboundOrders: models.InboundOrders{
					ID:             1,
					OrderDate:      "2023-10-01",
					OrderNumber:    12345,
					EmployeeID:     1,
					ProductBatchID: "1",
					WarehouseID:    1,
				},
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			inboundOrdersJSON: `{
					"id": 1,
					"order_date": "2023-10-01",
					"order_number": 12345,
					"employee_id": 1,
					"product_batch_id": "invalid",
					"warehouse_id": 1
				}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "422 - When passing a body with a valid json with missing parameters",
			inboundOrdersJSON: `{
					"id": 1
				}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute Order Date invalid",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			inboundOrdersJSON: `{
					"id": 1,
					"order_date": "2023-10-01",
					"order_number": 12345,
					"employee_id": 1,
					"product_batch_id": 1,
					"warehouse_id": 1
				}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:         1,
				statusCode:    http.StatusConflict,
				message:       "1062",
				inboundOrders: models.InboundOrders{},
			},
		},
		{
			name: "409 - When the repository raises a CannotAddOrUpdateChildRow error",
			inboundOrdersJSON: `{
					"id": 1,
					"order_date": "2023-10-01",
					"order_number": 12345,
					"employee_id": 1,
					"product_batch_id": 1,
					"warehouse_id": 1
				}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			wanted: wanted{
				calls:         1,
				statusCode:    http.StatusConflict,
				message:       "1452",
				inboundOrders: models.InboundOrders{},
			},
		},
		{
			name: "500 - When the repository returns an unexpected error",
			inboundOrdersJSON: `{
					"id": 1,
					"order_date": "2023-10-01",
					"order_number": 12345,
					"employee_id": 1,
					"product_batch_id": 1,
					"warehouse_id": 1
				}`,
			callErr: errors.New("Erro interno!"),
			wanted: wanted{
				calls:         1,
				statusCode:    http.StatusInternalServerError,
				message:       "Erro interno!",
				inboundOrders: models.InboundOrders{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewInboundOrdersServiceMock()
			ctl := controllers.NewInboundOrdersController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", strings.NewReader(tt.inboundOrdersJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On("Create", mock.AnythingOfType("models.InboundOrdersDTO")).Return(tt.wanted.inboundOrders, tt.callErr)
			ctl.Create(res, req)

			// Assert
			var decodedRes struct {
				Message string                                   `json:"message,omitempty"`
				Data    inboundordersctl.InboundOrdersAttributes `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			sv.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.inboundOrders.ID, decodedRes.Data.ID)
				require.Equal(t, tt.wanted.inboundOrders.EmployeeID, decodedRes.Data.EmployeeID)
				require.Equal(t, tt.wanted.inboundOrders.ProductBatchID, decodedRes.Data.ProductBatchID)
				require.Equal(t, tt.wanted.inboundOrders.WarehouseID, decodedRes.Data.WarehouseID)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
