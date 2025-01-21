package sectionsctl_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sectionsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	sectionsrp "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestReportPurchaseOrders(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name     string
		id       string
		sections []models.ProductReport
		callErr  error
		wanted   wanted
	}{
		{
			name: "200 - When sections are registered in the database and no id is passed",
			sections: []models.ProductReport{
				{SectionID: 1, SectionNumber: "123456789", ProductsCount: 1},
				{SectionID: 2, SectionNumber: "987654321", ProductsCount: 2},
				{SectionID: 3, SectionNumber: "111111111", ProductsCount: 3},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name: "200 - When sections are registered in the database and a valid id is passed",
			sections: []models.ProductReport{
				{SectionID: 2, SectionNumber: "987654321", ProductsCount: 2},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:     "200 - When no sections are registered in the database and no id is passed",
			id:       "2",
			sections: []models.ProductReport{},
			callErr:  nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:    "400 - When passing a non numeric id",
			id:      "S2",
			callErr: nil,
			wanted: wanted{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name:    "400 - When passing a negative id",
			id:      "-1",
			callErr: nil,
			wanted: wanted{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name:    "404 - When the respository raises a NotFound error",
			id:      "999",
			callErr: models.ErrSectionNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "section not found",
			},
		},
		{
			name:     "500 - When the repository returns an error",
			sections: []models.ProductReport{},
			callErr:  errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			sv := service.NewSectionsService(rp)
			ctl := controllers.NewSectionsController(sv)

			url := "/api/v1/sections/reportProducts"
			if tt.id != "" {
				url += fmt.Sprintf("?id=%s", tt.id)
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// Act
			rp.On("ReportProducts", mock.AnythingOfType("int")).Return(tt.sections, tt.callErr)
			ctl.ReportProducts(res, req)

			var decodedRes struct {
				Message string                              `json:"message,omitempty"`
				Data    []sectionsctl.ReportProductFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			rp.AssertNumberOfCalls(t, "ReportProducts", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			if len(tt.sections) > 0 {
				for i, buyer := range tt.sections {
					require.Equal(t, buyer.SectionID, decodedRes.Data[i].SectionID)
					require.Equal(t, buyer.SectionNumber, decodedRes.Data[i].SectionNumber)
					require.Equal(t, buyer.ProductsCount, decodedRes.Data[i].ProductsCount)
				}
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
