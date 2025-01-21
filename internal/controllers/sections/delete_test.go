package sectionsctl_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	sectionsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type wanted struct {
		calls      int
		statusCode int
		message    string
	}
	tests := []struct {
		name    string
		id      string
		callErr error
		wanted  wanted
	}{
		{
			name:    "204 - When the section is deleted successfully",
			id:      "1",
			callErr: nil,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNoContent,
			},
		},
		{
			name:    "400 - When passing a non numeric id",
			id:      "abc",
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
			name:    "404 - When the repository raises a SectionNotFound error",
			id:      "999",
			callErr: models.ErrSectionNotFound,
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusNotFound,
				message:    "section not found",
			},
		},
		{
			name:    "409 - When the repository raises a CannotDeleteOrUpdateParentRow error",
			id:      "1",
			callErr: &mysql.MySQLError{Number: mysqlerr.CodeCannotDeleteOrUpdateParentRow},
			wanted: wanted{
				calls:      1,
				statusCode: http.StatusConflict,
				message:    "1451",
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
			r.Delete("/api/v1/sections/{id}", ctl.Delete)
			url := "/api/v1/sections/" + string(tt.id)
			req := httptest.NewRequest(http.MethodDelete, url, nil)
			res := httptest.NewRecorder()

			// Act
			sv.On("Delete", mock.AnythingOfType("int")).Return(tt.callErr)
			r.ServeHTTP(res, req)
			ctl.Delete(res, req)

			var decodedRes struct {
				Message string                        `json:"message,omitempty"`
				Data    []sectionsctl.SectionFullJSON `json:"data"`
			}
			err := json.NewDecoder(res.Body).Decode(&decodedRes)

			// Assert
			sv.AssertNumberOfCalls(t, "Delete", tt.wanted.calls)
			require.NoError(t, err)
			require.Equal(t, tt.wanted.statusCode, res.Code)
			require.Contains(t, decodedRes.Message, tt.wanted.message)
		})
	}
}
