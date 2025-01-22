package sectionsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sectionsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		section    models.Section
	}
	tests := []struct {
		name    string
		id      string
		callErr error
		wanted  struct {
			calls      int
			statusCode int
			message    string
			section    models.Section
		}
	}{
		{
			name:    "200 - When sections the section is found",
			id:      "1",
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				section: models.Section{
					ID:                 1,
					SectionNumber:      "Section 1",
					CurrentTemperature: 22.5,
					MinimumTemperature: 18.0,
					CurrentCapacity:    50,
					MinimumCapacity:    20,
					MaximumCapacity:    100,
					WarehouseID:        1,
					ProductTypeID:      1,
				},
			},
		},
		{
			name:    "400 - When passing a non numeric id",
			id:      "abc",
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name:    "400 - When passing a negative id",
			id:      "-1",
			callErr: nil,
			wanted: wanted{
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
			name:    "500 - When the repository returns an error",
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
			// Arrange
			sv := service.NewSectionsServiceMock()
			ctl := controllers.NewSectionsController(sv)

			r := chi.NewRouter()
			r.Get("/api/v1/sections/{id}", ctl.GetByID)
			url := "/api/v1/sections/" + tt.id
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("GetByID", mock.AnythingOfType("int")).Return(tt.wanted.section, tt.callErr)
			r.ServeHTTP(res, req)

			var decodedRes struct {
				Message string                      `json:"message,omitempty"`
				Data    sectionsctl.SectionFullJSON `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "GetByID", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Equal(t, tt.wanted.section.ID, decodedRes.Data.ID)
			require.Equal(t, tt.wanted.section.SectionNumber, decodedRes.Data.SectionNumber)
			require.Equal(t, tt.wanted.section.CurrentTemperature, decodedRes.Data.CurrentTemperature)
			require.Equal(t, tt.wanted.section.MinimumTemperature, decodedRes.Data.MinimumTemperature)
			require.Equal(t, tt.wanted.section.CurrentCapacity, decodedRes.Data.CurrentCapacity)
			require.Equal(t, tt.wanted.section.MinimumCapacity, decodedRes.Data.MinimumCapacity)
			require.Equal(t, tt.wanted.section.MaximumCapacity, decodedRes.Data.MaximumCapacity)
			require.Equal(t, tt.wanted.section.WarehouseID, decodedRes.Data.WarehouseID)
			require.Equal(t, tt.wanted.section.ProductTypeID, decodedRes.Data.ProductTypeID)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
