package sectionsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sectionsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	sectionsrp "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
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
		section    models.Section
	}
	tests := []struct {
		name        string
		sectionJSON string
		callErr     error
		wanted      wanted
	}{
		{
			name: "201 - When the section is created successfully",
			sectionJSON: `{
				"section_number": "Section 1",
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"current_capacity": 50,
				"minimum_capacity": 10,
				"maximum_capacity": 100,
				"warehouse_id": 1,
				"product_type_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
				section: models.Section{
					ID:                 1,
					SectionNumber:      "Section 1",
					CurrentTemperature: 22.5,
					MinimumTemperature: 18.0,
					CurrentCapacity:    50,
					MinimumCapacity:    10,
					MaximumCapacity:    100,
					WarehouseID:        1,
					ProductTypeID:      1,
				},
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			sectionJSON: `{
				"section_number": "Section 1",
				"current_temperature": "22.5",
				"minimum_temperature": 18.0,
				"current_capacity": "50",
				"minimum_capacity": 10,
				"maximum_capacity": "100",
				"warehouse_id": 1,
				"product_type_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			sectionJSON: `{
				"section_number": "Section 1",
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"current_capacity": 50,
				"minimum_capacity": 10,
				"maximum_capacity": 100,
				"warehouse_id": 1,
				"product_type_id": 1
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				section:    models.Section{},
			},
		},
		{
			name: "409 - When the repository raises a CannotAddOrUpdateChildRow error",
			sectionJSON: `{
				"section_number": "Section 1",
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"current_capacity": 50,
				"minimum_capacity": 10,
				"maximum_capacity": 100,
				"warehouse_id": 1,
				"product_type_id": 1
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1452",
				section:    models.Section{},
			},
		},
		{
			name: "422 - When passing a body with a valid json with missing parameters",
			sectionJSON: `{
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"current_capacity": 50,
				"minimum_capacity": 10,
				"maximum_capacity": 100,
				"product_type_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute SectionNumber cannot be nil",
			},
		},
		{
			name: "422 - When passing a body with a valid json with empty parameters",
			sectionJSON: `{
				"section_number": "",
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"current_capacity": 0,
				"minimum_capacity": 10,
				"maximum_capacity": 0,
				"warehouse_id": 1,
				"product_type_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute SectionNumber cannot be empty",
			},
		},
		{
			name: "500 - When the repository returns an unexpected error",
			sectionJSON: `{
				"section_number": "Section 1",
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"current_capacity": 50,
				"minimum_capacity": 10,
				"maximum_capacity": 100,
				"warehouse_id": 1,
				"product_type_id": 1
			}`,
			callErr: errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				section:    models.Section{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			sv := service.NewSectionsService(rp)
			ctl := controllers.NewSectionsController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/sections", strings.NewReader(tt.sectionJSON))
			res := httptest.NewRecorder()

			// Act
			rp.On("Create", mock.AnythingOfType("models.SectionDTO")).Return(tt.wanted.section, tt.callErr)
			ctl.Create(res, req)

			var decodedRes struct {
				Message string                      `json:"message,omitempty"`
				Data    sectionsctl.SectionFullJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			// Assert
			rp.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.section.ID, decodedRes.Data.ID)
				require.Equal(t, tt.wanted.section.SectionNumber, decodedRes.Data.SectionNumber)
				require.Equal(t, tt.wanted.section.CurrentTemperature, decodedRes.Data.CurrentTemperature)
				require.Equal(t, tt.wanted.section.MinimumTemperature, decodedRes.Data.MinimumTemperature)
				require.Equal(t, tt.wanted.section.CurrentCapacity, decodedRes.Data.CurrentCapacity)
				require.Equal(t, tt.wanted.section.MinimumCapacity, decodedRes.Data.MinimumCapacity)
				require.Equal(t, tt.wanted.section.MaximumCapacity, decodedRes.Data.MaximumCapacity)
				require.Equal(t, tt.wanted.section.WarehouseID, decodedRes.Data.WarehouseID)
				require.Equal(t, tt.wanted.section.ProductTypeID, decodedRes.Data.ProductTypeID)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
