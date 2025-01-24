package purchaseordersctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	purchaseordersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/purchase_orders"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPurchaseOrders_Create(t *testing.T) {
	type expected struct {
		calls          int
		purchaseOrders models.PurchaseOrders
		statusCode     int
		message        string
	}
	tests := []struct {
		name               string
		purchaseOrdersJSON string
		callErr            error
		expected           expected
	}{
		{
			name: "201 - when successfully creating a purchase order",
			purchaseOrdersJSON: `
				{
					"order_number": "order#1a",
					"order_date": "2021-04-04",
					"tracking_code": "absc",
					"buyer_id": 1,
					"product_record_id": 1
				}`,
			callErr: nil,
			expected: expected{
				calls: 1,
				purchaseOrders: models.PurchaseOrders{
					OrderNumber:     "order#1a",
					OrderDate:       "2021-04-04",
					TrackingCode:    "absc",
					BuyerID:         1,
					ProductRecordID: 1,
				},
				statusCode: http.StatusCreated,
				message:    "Created",
			},
		},
		{
			name: "400 - when the request body is invalid",
			purchaseOrdersJSON: `
				{
					"order_number": "order#1a",
					"order_date": "2021-04-04",
					"tracking_code": "absc",
					"buyer_id": 1,
					"product_record_id": "s"
				}`,
			callErr: nil,
			expected: expected{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "400 - when the order date format is invalid",
			purchaseOrdersJSON: `
				{
					"order_number": "order#1a",
					"order_date": "04/04/2025",
					"tracking_code": "absc",
					"buyer_id": 1,
					"product_record_id": 1
				}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeIncorrectDateValue},
			expected: expected{
				calls:      1,
				statusCode: http.StatusBadRequest,
				message:    "1292",
			},
		},
		{
			name: "422 - when the order number is required",
			purchaseOrdersJSON: `
				{
					"order_date": "2021-04-04",
					"tracking_code": "absc",
					"buyer_id": 1,
					"product_record_id": 1
				}`,
			callErr: nil,
			expected: expected{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: order_number is required",
			},
		},
		{
			name: "409 - when the order number is duplicate",
			purchaseOrdersJSON: `
				{
					"order_number": "order#1a",
					"order_date": "2021-04-04",
					"tracking_code": "absc",
					"buyer_id": 1,
					"product_record_id": 1
				}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			expected: expected{
				calls:          1,
				purchaseOrders: models.PurchaseOrders{},
				statusCode:     http.StatusConflict,
				message:        "1062",
			},
		},
		{
			name: "409 - when the buyer ID does not exist",
			purchaseOrdersJSON: `
				{
					"order_number": "order#1a",
					"order_date": "2021-04-04",
					"tracking_code": "absc",
					"buyer_id": 1,
					"product_record_id": 9999
				}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			expected: expected{
				calls:          1,
				purchaseOrders: models.PurchaseOrders{},
				statusCode:     http.StatusConflict,
				message:        "1452",
			},
		},
		{
			name: "500 - when an internal server error occurs",
			purchaseOrdersJSON: `
				{
					"order_number": "",
					"order_date": "2021-04-04",
					"tracking_code": "absc",
					"buyer_id": 1,
					"product_record_id": 1
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
			purchaseOrders := tt.expected.purchaseOrders
			//Arrange
			sv := service.NewPurchaseOrdersServiceMock()
			ctl := controllers.NewPurchaseOrdersController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/purchaseOrders", strings.NewReader(tt.purchaseOrdersJSON))
			res := httptest.NewRecorder()
			//Act
			sv.On("Create", mock.MatchedBy(func(dto models.PurchaseOrdersDTO) bool {
				if tt.callErr != nil {
					return true
				}
				return (dto.OrderNumber == tt.expected.purchaseOrders.OrderNumber &&
					dto.OrderDate == tt.expected.purchaseOrders.OrderDate &&
					dto.TrackingCode == tt.expected.purchaseOrders.TrackingCode &&
					dto.BuyerID == tt.expected.purchaseOrders.BuyerID &&
					dto.ProductRecordID == tt.expected.purchaseOrders.ProductRecordID)
			})).Return(tt.expected.purchaseOrders, tt.callErr)
			ctl.Create(res, req)

			var decodedRes struct {
				Message string                               `json:"message,omitempty"`
				Data    purchaseordersctl.PurchaseOrdersJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			// Assert
			sv.AssertNumberOfCalls(t, "Create", tt.expected.calls)
			require.Equal(t, tt.expected.statusCode, res.Code)
			if tt.expected.statusCode == http.StatusCreated {
				require.Equal(t, purchaseOrders.OrderNumber, *decodedRes.Data.OrderNumber)
				require.Equal(t, purchaseOrders.OrderDate, *decodedRes.Data.OrderDate)
				require.Equal(t, purchaseOrders.TrackingCode, *decodedRes.Data.TrackingCode)
				require.Equal(t, purchaseOrders.BuyerID, *decodedRes.Data.BuyerID)
				require.Equal(t, purchaseOrders.ProductRecordID, *decodedRes.Data.ProductRecordID)
			}
			require.Contains(t, decodedRes.Message, tt.expected.message)

		})
	}
}
