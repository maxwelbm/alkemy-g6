package productsctl_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	productsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProducts_ReportRecords(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name          string
		id            string
		reportRecords []models.ProductReportRecords
		callErr       error
		wanted        wanted
	}{
		{
			name: "200 - Successfully retrieved all product report records",
			id:   "",
			reportRecords: []models.ProductReportRecords{
				{ProductID: 1, Description: "Product 1", RecordsCount: 1},
				{ProductID: 2, Description: "Product 2", RecordsCount: 1},
				{ProductID: 3, Description: "Product 3", RecordsCount: 1},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name: "200 - Successfully retrieved report records for a specific product by ID",
			id:   "1",
			reportRecords: []models.ProductReportRecords{
				{ProductID: 1, Description: "Product 1", RecordsCount: 1},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:          "200 - No report records found for a specific product ID",
			id:            "10",
			reportRecords: []models.ProductReportRecords{},
			callErr:       nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name: "400 - Invalid product ID format",
			id:   "invalid",
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name: "400 - Negative product ID value",
			id:   "-1",
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "404 - Product report record not found for the specified ID",
			id:      "10",
			callErr: models.ErrReportRecordNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "product report record not found",
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
			reportRecords := tt.reportRecords
			// Arrange
			sv := service.NewProductsServiceMock()
			ctl := controllers.NewProductsController(sv)

			url := "/api/v1/products/reportRecords"
			if tt.id != "" {
				url += fmt.Sprintf("?id=%s", tt.id)
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("ReportRecords", mock.AnythingOfType("int")).Return(reportRecords, tt.callErr)
			ctl.ReportRecords(res, req)

			var decodedRes struct {
				Message string                             `json:"message,omitempty"`
				Data    []productsctl.ReportRecordFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "ReportRecords", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			if len(tt.reportRecords) > 0 {
				for i, value := range reportRecords {
					require.Equal(t, decodedRes.Data[i].ProductID, value.ProductID)
					require.Equal(t, decodedRes.Data[i].Description, value.Description)
					require.Equal(t, decodedRes.Data[i].RecordsCount, value.RecordsCount)
				}
			}

			require.Contains(t, decodedRes.Message, tt.wanted.message)

		})
	}

}
