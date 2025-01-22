package localitiesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	localitiesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/localities"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		locality   models.Locality
	}
	tests := []struct {
		name         string
		localityJSON string
		callErr      error
		wanted       wanted
	}{
		{
			name:         "201 - When the locality is created successfully",
			localityJSON: `{"locality_name": "São Paulo", "province_name": "São Paulo", "country_name": "Brazil"}`,
			callErr:      nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
				locality:   models.Locality{ID: 1, LocalityName: "São Paulo", ProvinceName: "São Paulo", CountryName: "Brazil"},
			},
		},
		{
			name:         "400 - When passing a body with invalid json",
			localityJSON: `{"locality_name": 1, "province_name": 2, "country_name": 3}`,
			callErr:      nil,
			wanted: wanted{
				calls:      0,
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name:         "409 - When the repository raises a DuplicateEntry error",
			localityJSON: `{"locality_name": "São Paulo", "province_name": "São Paulo", "country_name": "Brazil"}`,
			callErr:      &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				locality:   models.Locality{},
			},
		},
		{
			name:         "422 - When passing a body with a valid json with missing parameters",
			localityJSON: `{"province_name": "São Paulo", "country_name": "Brazil"}`,
			callErr:      nil,
			wanted: wanted{
				calls:      0,
				statusCode: http.StatusUnprocessableEntity,
				message:    "locality_name cannot be nil",
			},
		},
		{
			name:         "422 - When passing a body with a valid json with empty parameters",
			localityJSON: `{"locality_name": "", "province_name": "São Paulo", "country_name": "Brazil"}`,
			callErr:      nil,
			wanted: wanted{
				calls:      0,
				statusCode: http.StatusUnprocessableEntity,
				message:    "locality_name cannot be empty",
			},
		},
		{
			name:         "500 - When the repository returns an unexpected error",
			localityJSON: `{"locality_name": "São Paulo", "province_name": "São Paulo", "country_name": "Brazil"}`,
			callErr:      errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				locality:   models.Locality{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewLocalitiesServiceMock()
			ctl := controllers.NewLocalityController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/localities", strings.NewReader(tt.localityJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On("Create", mock.AnythingOfType("models.LocalityDTO")).Return(tt.wanted.locality, tt.callErr)
			ctl.Create(res, req)

			// Assert
			var decodedRes struct {
				Message string                        `json:"message,omitempty"`
				Data    localitiesctl.FullLocalitySON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			sv.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.locality.LocalityName, decodedRes.Data.LocalityName)
				require.Equal(t, tt.wanted.locality.ProvinceName, decodedRes.Data.ProvinceName)
				require.Equal(t, tt.wanted.locality.CountryName, decodedRes.Data.CountryName)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
