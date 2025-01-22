package sectionsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sectionsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name     string
		sections []models.Section
		callErr  error
		wanted   wanted
	}{
		{
			name: "200 - When sections are registered in the database",
			sections: []models.Section{
				{ID: 1, SectionNumber: "Section 1", CurrentTemperature: 22.5, MinimumTemperature: 18.0, CurrentCapacity: 50, MinimumCapacity: 20, MaximumCapacity: 100, WarehouseID: 1, ProductTypeID: 1},
				{ID: 2, SectionNumber: "Section 2", CurrentTemperature: 23.0, MinimumTemperature: 18.5, CurrentCapacity: 60, MinimumCapacity: 25, MaximumCapacity: 110, WarehouseID: 2, ProductTypeID: 2},
				{ID: 3, SectionNumber: "Section 3", CurrentTemperature: 21.0, MinimumTemperature: 17.0, CurrentCapacity: 70, MinimumCapacity: 30, MaximumCapacity: 120, WarehouseID: 3, ProductTypeID: 3},
			},
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:     "200 - When no sections are registered in the database",
			sections: []models.Section{},
			callErr:  nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
			},
		},
		{
			name:     "500 - When the repository returns an error",
			sections: []models.Section{},
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
			sv := service.NewSectionsServiceMock()
			ctl := controllers.NewSectionsController(sv)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("GetAll").Return(tt.sections, tt.callErr)
			ctl.GetAll(res, req)

			var decodedRes struct {
				Message string                        `json:"message,omitempty"`
				Data    []sectionsctl.SectionFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "GetAll", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			if len(tt.sections) > 0 {
				for i, section := range tt.sections {
					require.Equal(t, section.ID, decodedRes.Data[i].ID)
					require.Equal(t, section.SectionNumber, decodedRes.Data[i].SectionNumber)
					require.Equal(t, section.CurrentTemperature, decodedRes.Data[i].CurrentTemperature)
					require.Equal(t, section.MinimumTemperature, decodedRes.Data[i].MinimumTemperature)
					require.Equal(t, section.CurrentCapacity, decodedRes.Data[i].CurrentCapacity)
					require.Equal(t, section.MinimumCapacity, decodedRes.Data[i].MinimumCapacity)
					require.Equal(t, section.MaximumCapacity, decodedRes.Data[i].MaximumCapacity)
					require.Equal(t, section.WarehouseID, decodedRes.Data[i].WarehouseID)
					require.Equal(t, section.ProductTypeID, decodedRes.Data[i].ProductTypeID)
				}
			}
			if tt.wanted.message != "" {
				require.Contains(t, decodedRes.Message, tt.wanted.message)
			}
		})
	}
}
