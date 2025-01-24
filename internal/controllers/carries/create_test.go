package carriesctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	carriesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/carries"
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
		carrie     models.Carry
	}
	tests := []struct {
		name       string
		carrieJSON string
		callErr    error
		wanted     wanted
	}{
		{
			name: "201 - When the carrie is created successfully",
			carrieJSON: `{
				"cid": "123456789",
				"company_name": "Carry 1",
				"address": "Rua legal",
				"phone_number": "123456789",
				"locality_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusCreated,
				message:    "Created",
				carrie: models.Carry{
					ID:          1,
					CID:         "123456789",
					CompanyName: "Carry 1",
					Address:     "Rua legal",
					PhoneNumber: "123456789",
					LocalityID:  1,
				},
			},
		},
		{
			name: "400 - When passing a body with invalid json",
			carrieJSON: `{
				"cid": 123456789,
				"company_name": "Carry 1",
				"address": "Rua legal",
				"phone_number": "123456789",
				"locality_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusBadRequest,
				message:    "json: cannot unmarshal",
			},
		},
		{
			name: "422 - When passing a body with a valid json with missing parameters",
			carrieJSON: `{
				"cid": "123456789",
				"address": "Rua legal",
				"phone_number": "123456789",
				"locality_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "company_name cannot be nil",
			},
		},
		{
			name: "422 - When passing a body with a valid json with empty parameters",
			carrieJSON: `{
				"cid": "123456789",
				"company_name": "",
				"address": "Rua legal",
				"phone_number": "123456789",
				"locality_id": 1
			}`,
			callErr: nil,
			wanted: wanted{
				statusCode: http.StatusUnprocessableEntity,
				message:    "company_name cannot be empty",
			},
		},
		{
			name: "409 - When the repository raises a DuplicateEntry error",
			carrieJSON: `{
				"cid": "123456789",
				"company_name": "Carry 1",
				"address": "Rua legal",
				"phone_number": "123456789",
				"locality_id": 1
			}`,
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1062",
				carrie:     models.Carry{},
			},
		},
		{
			name: "500 - When the repository returns an unexpected error",
			carrieJSON: `{
				"cid": "123456789",
				"company_name": "Carry 1",
				"address": "Rua legal",
				"phone_number": "123456789",
				"locality_id": 1
			}`,
			callErr: errors.New("internal error"),
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusInternalServerError,
				message:    "internal error",
				carrie:     models.Carry{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			sv := service.NewCarriesServiceMock()
			ctl := controllers.NewCarriesController(sv)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/carries", strings.NewReader(tt.carrieJSON))
			res := httptest.NewRecorder()

			// Act
			sv.On("Create", mock.MatchedBy(func(dto models.CarryDTO) bool {
				if tt.callErr != nil {
					return true
				}
				return (*dto.CID == tt.wanted.carrie.CID &&
					*dto.CompanyName == tt.wanted.carrie.CompanyName &&
					*dto.Address == tt.wanted.carrie.Address &&
					*dto.PhoneNumber == tt.wanted.carrie.PhoneNumber &&
					*dto.LocalityID == tt.wanted.carrie.LocalityID)
			})).Return(tt.wanted.carrie, tt.callErr)
			ctl.Create(res, req)

			// Assert
			var decodedRes struct {
				Message string                   `json:"message,omitempty"`
				Data    carriesctl.FullCarryJSON `json:"data,omitempty"`
			}
			require.NoError(t, json.NewDecoder(res.Body).Decode(&decodedRes))

			sv.AssertNumberOfCalls(t, "Create", tt.wanted.calls)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
			if tt.wanted.statusCode == http.StatusCreated {
				require.Equal(t, tt.wanted.carrie.CID, decodedRes.Data.CID)
				require.Equal(t, tt.wanted.carrie.CompanyName, decodedRes.Data.CompanyName)
				require.Equal(t, tt.wanted.carrie.Address, decodedRes.Data.Address)
				require.Equal(t, tt.wanted.carrie.PhoneNumber, decodedRes.Data.PhoneNumber)
				require.Equal(t, tt.wanted.carrie.LocalityID, decodedRes.Data.LocalityID)
			}
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
