package sectionsctl

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Create handles the creation of a new section.
// @Summary Create a new section
// @Description Create a new section with the provided JSON payload
// @Tags sections
// @Accept json
// @Produce json
// @Param section body NewSectionReqJSON true "New Section JSON"
// @Success 201 {object} models.SectionDTO "Created"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/sections [post]
func (ctl *SectionsController) Create(w http.ResponseWriter, r *http.Request) {
	var secReqJSON NewSectionReqJSON
	if err := json.NewDecoder(r.Body).Decode(&secReqJSON); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := secReqJSON.validateCreate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	secDTO := models.SectionDTO{
		SectionNumber:      secReqJSON.SectionNumber,
		CurrentTemperature: secReqJSON.CurrentTemperature,
		MinimumTemperature: secReqJSON.MinimumTemperature,
		CurrentCapacity:    secReqJSON.CurrentCapacity,
		MinimumCapacity:    secReqJSON.MinimumCapacity,
		MaximumCapacity:    secReqJSON.MaximumCapacity,
		WarehouseID:        secReqJSON.WarehouseID,
		ProductTypeID:      secReqJSON.ProductTypeID,
	}

	newSection, err := ctl.sv.Create(secDTO)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	data := SectionFullJSON{
		ID:                 newSection.ID,
		SectionNumber:      newSection.SectionNumber,
		CurrentTemperature: newSection.CurrentTemperature,
		MinimumTemperature: newSection.MinimumTemperature,
		CurrentCapacity:    newSection.CurrentCapacity,
		MinimumCapacity:    newSection.MinimumCapacity,
		MaximumCapacity:    newSection.MaximumCapacity,
		WarehouseID:        newSection.WarehouseID,
		ProductTypeID:      newSection.ProductTypeID,
	}

	res := SectionResJSON{
		Message: "Created",
		Data:    data,
	}

	response.JSON(w, http.StatusCreated, res)
}
