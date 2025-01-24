package productbatchesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	productbatchesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/product_batches"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProductBatch_Create(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		prodBatch  models.ProductBatches
	}
	tests := []struct {
		name          string
		prodBatchJSON string
		callErr       error
		wanted        wanted
	}{
		{
			name: "201 - When the section is created successfully",
			prodBatchJSON: `{
				"batch_number": "PB01",
				"initial_quantity" : 100,
				"current_quantity" : 50,
				"current_temperature": 20.0,
				"minimum_temperature":-5.0,
				"due_date": "2022-04-04",
				"manufacturing_date": "2020-04-04",
				"manufacturing_hour": "08:00:00",
				"product_id": 1,
				"section_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
				prodBatch: models.ProductBatches{
					ID:                 1,
					BatchNumber:        "PB01",
					InitialQuantity:    100,
					CurrentQuantity:    50,
					CurrentTemperature: 20.0,
					MinimumTemperature: -5.0,
					DueDate:            "2022-04-04",
					ManufacturingDate:  "2020-04-04",
					ManufacturingHour:  "08:00:00",
					ProductID:          1,
					SectionID:          1,
				},
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			prodBatchJSON: `{
					"batch_number": 1,
					"initial_quantity" : 100,
					"current_quantity" : 50,
					"current_temperature": 20.0,
					"minimum_temperature":-5.0,
					"due_date": "2022-04-04",
					"manufacturing_date": "2020-04-04",
					"manufacturing_hour": "08:00:00",
					"product_id": 1,
					"section_id": 1
				}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			prodBatchJSON: `{
					"batch_number": "PB01",
					"initial_quantity" : 100,
					"current_quantity" : 50,
					"current_temperature": 20.0,
					"minimum_temperature":-5.0,
					"due_date": "2022-04-04",
					"manufacturing_date": "2020-04-04",
					"manufacturing_hour": "08:00:00",
					"product_id": 1,
					"section_id": 1
				}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				prodBatch:  models.ProductBatches{},
			},
		},
		{
			name: "409 - When the repository raises a CannotAddOrUpdateChildRow error",
			prodBatchJSON: `{
					"batch_number": "PB01",
					"initial_quantity" : 100,
					"current_quantity" : 50,
					"current_temperature": 20.0,
					"minimum_temperature":-5.0,
					"due_date": "2022-04-04",
					"manufacturing_date": "2020-04-04",
					"manufacturing_hour": "08:00:00",
					"product_id": 1,
					"section_id": 1
				}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1452",
				prodBatch:  models.ProductBatches{},
			},
		},
		{
			name: "422 - When passing a body with a valid json with missing parameters",
			prodBatchJSON: `{
					"batch_number": "PB01",
					"current_quantity" : 50,
					"current_temperature": 20.0,
					"minimum_temperature":-5.0,
					"due_date": "2022-04-04",
					"manufacturing_date": "2020-04-04",
					"manufacturing_hour": "08:00:00",
					"product_id": 1,
					"section_id": 1
				}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute InitialQuantity cannot be nil",
			},
		},
		{
			name: "422 - When passing a body with a valid json with empty parameters",
			prodBatchJSON: `{
					"batch_number": "",
					"initial_quantity" : 100,
					"current_quantity" : 50,
					"current_temperature": 20.0,
					"minimum_temperature":-5.0,
					"due_date": "2022-04-04",
					"manufacturing_date": "2020-04-04",
					"manufacturing_hour": "08:00:00",
					"product_id": 1,
					"section_id": 1
				}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute BatchNumber cannot be empty",
			},
		},
		{
			name: "500 - When the repository returns an unexpected error",
			prodBatchJSON: `{
				"batch_number": "PB01",
				"initial_quantity" : 100,
				"current_quantity" : 50,
				"current_temperature": 20.0,
				"minimum_temperature":-5.0,
				"due_date": "2022-04-04",
				"manufacturing_date": "2020-04-04",
				"manufacturing_hour": "08:00:00",
				"product_id": 1,
				"section_id": 1
			}`,
			callErr: errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				prodBatch:  models.ProductBatches{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewProductBatchesServiceMock()
			ctl := controllers.NewProductBatchesController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/productBatches", strings.NewReader(tt.prodBatchJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On("Create", mock.MatchedBy(func(dto models.ProductBatchesDTO) bool {
				if tt.callErr != nil {
					return true
				}
				return (dto.BatchNumber == tt.wanted.prodBatch.BatchNumber &&
					dto.InitialQuantity == tt.wanted.prodBatch.InitialQuantity &&
					dto.CurrentQuantity == tt.wanted.prodBatch.CurrentQuantity &&
					dto.CurrentTemperature == tt.wanted.prodBatch.CurrentTemperature &&
					dto.MinimumTemperature == tt.wanted.prodBatch.MinimumTemperature &&
					dto.DueDate == tt.wanted.prodBatch.DueDate &&
					dto.ManufacturingDate == tt.wanted.prodBatch.ManufacturingDate &&
					dto.ManufacturingHour == tt.wanted.prodBatch.ManufacturingHour &&
					dto.SectionID == tt.wanted.prodBatch.SectionID &&
					dto.ProductID == tt.wanted.prodBatch.ProductID)
			})).Return(tt.wanted.prodBatch, tt.callErr)
			ctl.Create(res, req)

			var decodedRes struct {
				Message string                                 `json:"message,omitempty"`
				Data    productbatchesctl.ProductBatchFullJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			// Assert
			sv.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.prodBatch.ID, decodedRes.Data.ID)
				require.Equal(t, tt.wanted.prodBatch.BatchNumber, decodedRes.Data.BatchNumber)
				require.Equal(t, tt.wanted.prodBatch.InitialQuantity, decodedRes.Data.InitialQuantity)
				require.Equal(t, tt.wanted.prodBatch.CurrentQuantity, decodedRes.Data.CurrentQuantity)
				require.Equal(t, tt.wanted.prodBatch.CurrentTemperature, decodedRes.Data.CurrentTemperature)
				require.Equal(t, tt.wanted.prodBatch.MinimumTemperature, decodedRes.Data.MinimumTemperature)
				require.Equal(t, tt.wanted.prodBatch.DueDate, decodedRes.Data.DueDate)
				require.Equal(t, tt.wanted.prodBatch.ManufacturingDate, decodedRes.Data.ManufacturingDate)
				require.Equal(t, tt.wanted.prodBatch.ManufacturingHour, decodedRes.Data.ManufacturingHour)
				require.Equal(t, tt.wanted.prodBatch.SectionID, decodedRes.Data.SectionID)
				require.Equal(t, tt.wanted.prodBatch.ProductID, decodedRes.Data.ProductID)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
