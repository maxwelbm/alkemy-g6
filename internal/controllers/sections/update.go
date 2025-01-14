package sectionsctl

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// @Summary Update a section
// @Description Update a section by ID
// @Tags sections
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Param section body NewSectionReqJSON true "New Section JSON"
// @Success 200 {object} models.SectionDTO "Updated"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/sections/{id} [patch]
func (c *SectionsController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var secReqJson NewSectionReqJSON
	if err = json.NewDecoder(r.Body).Decode(&secReqJson); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = secReqJson.validateUpdate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	secDTO := models.SectionDTO{
		SectionNumber:      secReqJson.SectionNumber,
		CurrentTemperature: secReqJson.CurrentTemperature,
		MinimumTemperature: secReqJson.MinimumTemperature,
		CurrentCapacity:    secReqJson.CurrentCapacity,
		MinimumCapacity:    secReqJson.MinimumCapacity,
		MaximumCapacity:    secReqJson.MaximumCapacity,
		WarehouseID:        secReqJson.WarehouseID,
		ProductTypeID:      secReqJson.ProductTypeID,
	}

	updateSection, err := c.sv.Update(id, secDTO)

	if err != nil {
		// Handle if section not found
		if errors.Is(err, models.ErrSectionNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		// Handle no changes made
		if errors.Is(err, models.ErrorNoChangesMade) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		// Handle MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// Handle other internal server errors
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := SectionFullJSON{
		ID:                 updateSection.ID,
		SectionNumber:      updateSection.SectionNumber,
		CurrentTemperature: updateSection.CurrentTemperature,
		MinimumTemperature: updateSection.MinimumTemperature,
		CurrentCapacity:    updateSection.CurrentCapacity,
		MinimumCapacity:    updateSection.MinimumCapacity,
		MaximumCapacity:    updateSection.MaximumCapacity,
		WarehouseID:        updateSection.WarehouseID,
		ProductTypeID:      updateSection.ProductTypeID,
	}

	res := SectionResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}
