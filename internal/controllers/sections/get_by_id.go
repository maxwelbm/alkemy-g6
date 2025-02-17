package sectionsctl

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetByID handles the HTTP request to retrieve a section by its ID.
// @Summary Get a section by ID
// @Description Get a section by ID
// @Tags sections
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Success 200 {object} SectionFullJSON
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Router /api/v1/sections/{id} [get]
func (ctl *SectionsController) GetByID(w http.ResponseWriter, r *http.Request) {
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

	// Get the section from the service
	section, err := ctl.sv.GetByID(id)
	if err != nil {
		// If the section is not found, return a 404 Not Found error
		if errors.Is(err, models.ErrSectionNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

			return
		}

		// If the section is not found, return a 500 Internal Server Error error
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	// Map the section to the response
	data := SectionFullJSON{
		ID:                 section.ID,
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseID:        section.WarehouseID,
		ProductTypeID:      section.ProductTypeID,
	}

	// Return the response
	res := SectionResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusOK, res))
}
