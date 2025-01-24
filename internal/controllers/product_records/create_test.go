package productrecordsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	productrecordsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/product_records"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type expected struct {
		calls         int
		statusCode    int
		message       string
		productRecord models.ProductRecord
	}
	tests := []struct {
		name              string
		productRecordJSON string
		callErr           error
		expected          expected
	}{
		{
			name:              "201 - Successfully created a productRecord",
			productRecordJSON: `{"last_update_date":"2021-09-01","purchase_price":10.5,"sale_price":15.5,"product_id":1}`,
			callErr:           nil,
			expected: expected{
				calls:         1,
				statusCode:    http.StatusCreated,
				productRecord: models.ProductRecord{ID: 1, LastUpdateDate: "2021-09-01", PurchasePrice: 10.5, SalePrice: 15.5, ProductID: 1},
			},
		},
		{
			name:              "400 - Bad Request error when trying to create a productRecord",
			productRecordJSON: `{"last_update_date":"2021-09-01","purchase_price":"10.5","sale_price":15.5,"product_id":1}`,
			callErr:           nil,
			expected: expected{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name:              "409 - Conflict error when trying to create an productRecords with an inexistent productID",
			productRecordJSON: `{"last_update_date":"2021-09-01","purchase_price":10.5,"sale_price":15.5,"product_id":10}`,
			callErr:           &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			expected: expected{
				calls:         1,
				statusCode:    http.StatusConflict,
				message:       "1452",
				productRecord: models.ProductRecord{},
			},
		},
		{
			name:              "422 - Conflict error when trying to create an productRecords with empty parameters",
			productRecordJSON: `{"last_update_date":"","purchase_price":10.5,"sale_price":15.5,"product_id":1}`,
			callErr:           nil,
			expected: expected{
				calls:         0,
				statusCode:    http.StatusUnprocessableEntity,
				message:       "error: attribute LastUpdateDate cannot be empty",
				productRecord: models.ProductRecord{},
			},
		},
		{
			name:              "422 - Conflict error when trying to create an productRecords with missing parameters",
			productRecordJSON: `{"last_update_date":"2021-09-01","purchase_price":10.5,"sale_price":15.5}`,
			callErr:           nil,
			expected: expected{
				calls:         0,
				statusCode:    http.StatusUnprocessableEntity,
				message:       "error: attribute ProductID cannot be nil",
				productRecord: models.ProductRecord{},
			},
		},
		{
			name:              "500 - Internal Server Error when trying to retrieve the list of productRecordss",
			productRecordJSON: `{"last_update_date":"2021-09-01","purchase_price":10.5,"sale_price":15.5,"product_id":1}`,
			callErr:           errors.New("internal error"),
			expected: expected{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := service.NewProductRecordsServiceMock()
			ctl := controllers.NewProductRecordsController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/productRecords", strings.NewReader(tt.productRecordJSON))
			res := httptest.NewRecorder()

			sv.On("Create", mock.MatchedBy(func(dto models.ProductRecordDTO) bool {
                if tt.callErr != nil {
                    return true
                }
                return (dto.LastUpdateDate == tt.expected.productRecord.LastUpdateDate &&
                    dto.ProductID == tt.expected.productRecord.ProductID &&
                    dto.PurchasePrice == tt.expected.productRecord.PurchasePrice &&
					dto.SalePrice == tt.expected.productRecord.SalePrice)
            })).Return(tt.expected.productRecord, tt.callErr)

			ctl.Create(res, req)

			var decodedRes struct {
				Message string                                  `json:"message,omitempty"`
				Data    productrecordsctl.FullProductRecordJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			sv.AssertNumberOfCalls(t, "Create", tt.expected.calls)
			require.NoError(t, err)
			require.Equal(t, tt.expected.statusCode, res.Code)
			if tt.expected.statusCode == http.StatusOK {
				require.Equal(t, tt.expected.productRecord.ID, decodedRes.Data.ID)
				require.Equal(t, tt.expected.productRecord.LastUpdateDate, decodedRes.Data.LastUpdateDate)
				require.Equal(t, tt.expected.productRecord.PurchasePrice, decodedRes.Data.PurchasePrice)
				require.Equal(t, tt.expected.productRecord.SalePrice, decodedRes.Data.SalePrice)
				require.Equal(t, tt.expected.productRecord.ProductID, decodedRes.Data.ProductID)
			}
			require.Contains(t, decodedRes.Message, tt.expected.message)
		})
	}
}
