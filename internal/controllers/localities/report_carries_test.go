package localitiesctl_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	localitiesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/localities"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestReportCarries(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name    string
		id      string
		reports []models.LocalityCarriesReport
		callErr error
		wanted  wanted
	}{
		{
			name: "200 - When localities are registered in the database and no id is passed",
			reports: []models.LocalityCarriesReport{
				{ID: 1, LocalityName: "Locality 1", CarriesCount: 1},
				{ID: 2, LocalityName: "Locality 2", CarriesCount: 2},
				{ID: 3, LocalityName: "Locality 3", CarriesCount: 3},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name: "200 - When localities are registered in the database and a valid id is passed",
			reports: []models.LocalityCarriesReport{
				{ID: 2, LocalityName: "Locality 2", CarriesCount: 2},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:    "200 - When no localities are registered in the database and no id is passed",
			id:      "2",
			reports: []models.LocalityCarriesReport{},
			callErr: nil,
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
			callErr: models.ErrLocalityNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "locality not found",
			},
		},
		{
			name:    "500 - When the repository returns an error",
			reports: []models.LocalityCarriesReport{},
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
			// Arrange
			sv := service.NewLocalitiesServiceMock()
			ctl := controllers.NewLocalityController(sv)

			url := "/api/v1/localities/reportCarries"
			if tt.id != "" {
				url += fmt.Sprintf("?id=%s", tt.id)
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("ReportCarries", mock.AnythingOfType("int")).Return(tt.reports, tt.callErr)
			ctl.ReportCarries(res, req)

			var decodedRes struct {
				Message string                                    `json:"message,omitempty"`
				Data    []localitiesctl.LocalityCarriesReportJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "ReportCarries", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			if len(tt.reports) > 0 {
				for i, locality := range tt.reports {
					require.Equal(t, locality.ID, decodedRes.Data[i].ID)
					require.Equal(t, locality.LocalityName, decodedRes.Data[i].LocalityName)
					require.Equal(t, locality.CarriesCount, decodedRes.Data[i].CarriesCount)
				}
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
