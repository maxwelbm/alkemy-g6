package sectionsctl

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Delete handles the deletion of a section by ID.
// @Summary Delete a section
// @Description Delete a section by ID
// @Tags sections
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Success 204
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/sections/{id} [delete]
func (ctl *SectionsController) Delete(w http.ResponseWriter, r *http.Request) {
	// Parse the section ID from the URL parameter
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	// If the ID is invalid, return a 400 Bad Request error
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}
	// If the ID is less than 1, return a 400 Bad Request error
	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))

		return
	}

	err = ctl.sv.Delete(id)

	if err != nil {
		// Handle if section not found
		if errors.Is(err, models.ErrSectionNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

			return
		}
		// Handle MySQL conflict dependencies
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotDeleteOrUpdateParentRow {
			response.Error(w, http.StatusConflict, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusConflict, err.Error()))

			return
		}
		// Handle other internal server errors
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	response.JSON(w, http.StatusNoContent, nil)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusNoContent, nil))
}
