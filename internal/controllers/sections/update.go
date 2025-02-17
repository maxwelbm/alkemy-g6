package sectionsctl

import (
	"encoding/json"
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

type UpdateSectionReqJSON struct {
	SectionNumber      *string  `json:"section_number"`
	CurrentTemperature *float64 `json:"current_temperature"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
	CurrentCapacity    *int     `json:"current_capacity"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MaximumCapacity    *int     `json:"maximum_capacity"`
	WarehouseID        *int     `json:"warehouse_id"`
	ProductTypeID      *int     `json:"product_type_id"`
}

//nolint:gocyclo
func (sec *UpdateSectionReqJSON) validate() (err error) {
	var validationErrors []string

	// Check for nil pointers and collect their errors
	if sec.SectionNumber != nil && *sec.SectionNumber == "" {
		validationErrors = append(validationErrors, "error: attribute SectionNumber cannot be empty")
	}

	if sec.CurrentCapacity != nil {
		if *sec.CurrentCapacity < 0 {
			validationErrors = append(validationErrors, "error: attribute CurrentCapacity cannot be negative")
		}

		if *sec.CurrentCapacity == 0 {
			validationErrors = append(validationErrors, "error: attribute CurrentCapacity cannot be 0")
		}
	}

	if sec.MinimumCapacity != nil {
		if *sec.MinimumCapacity < 0 {
			validationErrors = append(validationErrors, "error: attribute MinimumCapacity cannot be negative")
		}

		if *sec.MinimumCapacity == 0 {
			validationErrors = append(validationErrors, "error: attribute MinimumCapacity cannot be 0")
		}
	}

	if sec.MaximumCapacity != nil {
		if *sec.MaximumCapacity < 0 {
			validationErrors = append(validationErrors, "error: attribute MaximumCapacity cannot be negative")
		}

		if *sec.MaximumCapacity == 0 {
			validationErrors = append(validationErrors, "error: attribute MaximumCapacity cannot be 0")
		}
	}

	if sec.WarehouseID != nil {
		if *sec.WarehouseID < 0 {
			validationErrors = append(validationErrors, "error: attribute WarehouseID cannot be negative")
		}

		if *sec.WarehouseID == 0 {
			validationErrors = append(validationErrors, "error: attribute WarehouseID cannot be 0")
		}
	}

	if sec.ProductTypeID != nil {
		if *sec.ProductTypeID < 0 {
			validationErrors = append(validationErrors, "error: attribute ProductTypeID cannot be negative")
		}

		if *sec.ProductTypeID == 0 {
			validationErrors = append(validationErrors, "error: attribute ProductTypeID cannot be 0")
		}
	}

	// Combine all errors before returning
	if len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, validationErrors...)

		err = fmt.Errorf("validation errors: %v", allErrors)
	}

	return err
}

// Update - Updates an existing section
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
func (ctl *SectionsController) Update(w http.ResponseWriter, r *http.Request) {
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
	// Parse the request body
	var secReqJSON UpdateSectionReqJSON
	if err = json.NewDecoder(r.Body).Decode(&secReqJSON); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}
	// Validate the request
	if err = secReqJSON.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusUnprocessableEntity, err.Error()))

		return
	}

	// Update the section
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

	updateSection, err := ctl.sv.Update(id, secDTO)

	//	Handle update errors
	if err != nil {
		// Handle if section not found
		if errors.Is(err, models.ErrSectionNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

			return
		}
		// Handle MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok &&
			(mysqlErr.Number == mysqlerr.CodeDuplicateEntry ||
				mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow) {
			response.Error(w, http.StatusConflict, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusConflict, err.Error()))

			return
		}
		// Handle other internal server errors
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

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
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}
