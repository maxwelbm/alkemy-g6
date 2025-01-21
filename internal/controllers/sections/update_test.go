package sectionsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	sectionsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
		section    models.Section
	}
	tests := []struct {
		name        string
		id          string
		sectionJSON string
		callErr     error
		wanted      wanted
	}{
		{
			name: "200 - When the section is updated successfully with all fields",
			id:   "1",
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
				statusCode: http.StatusOK,
				message:    "OK",
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
			name: "200 - When the section is updated successfully with missing fields",
			id:   "1",
			sectionJSON: `{
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"minimum_capacity": 10,
				"maximum_capacity": 100,
				"warehouse_id": 1,
				"product_type_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusOK,
				message:    "OK",
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
			name: "400 - When passing a non numeric id",
			id:   "abc",
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
				statusCode: http.StatusBadRequest,
				message:    "strconv.Atoi",
			},
		},
		{
			name: "400 - When passing a negative id",
			id:   "-1",
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
				statusCode: http.StatusBadRequest,
				message:    "Bad Request",
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			id:   "1",
			sectionJSON: `{
				"section_number": 1,
				"current_temperature": 22.5,
				"minimum_temperature": 18.0,
				"current_capacity": 50,
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
			name: "404 - When the repository raises a SectionNotFound error",
			id:   "999",
			sectionJSON: `{
				"section_number": "Section 1",
				"current_temperature": 22.5,
				"minimum_temperature": 18.0
			}`,
			callErr: models.ErrSectionNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "section not found",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			id:   "1",
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
			},
		},
		{
			name: "409 - When the repository raises a CannotAddOrUpdateChildRow error",
			id:   "1",
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
			},
		},
		{
			name: "422 - When passing a body with empty fields",
			id:   "1",
			sectionJSON: `{
				"section_number": "",
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
				statusCode: http.StatusUnprocessableEntity,
				message:    "error: attribute SectionNumber cannot be empty",
			},
		},
		{
			name: "500 - When the repository returns an error",
			id:   "1",
			sectionJSON: `{
				"section_number": "Section 1",
				"current_temperature": 22.5
			}`,
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
			ctl := sectionsctl.NewSectionsController(sv)

			r := chi.NewRouter()
			r.Patch("/api/v1/sections/{id}", ctl.Update)
			url := "/api/v1/sections/" + string(tt.id)
			req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(tt.sectionJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On(
				"Update",
				mock.AnythingOfType("int"),
				mock.AnythingOfType("models.SectionDTO"),
			).Return(tt.wanted.section, tt.callErr)
			r.ServeHTTP(res, req)

			// Assert
			var decodedRes struct {
				Message string         `json:"message,omitempty"`
				Data    models.Section `json:"data,omitempty"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)
			require.NoError(t, err)

			sv.AssertNumberOfCalls(t, "Update", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.callErr != nil {
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
		})
	}
}
