package sectionsctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type NewSectionReqJSON struct {
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
func (sec *NewSectionReqJSON) validate() (err error) {
	var validationErrors, nilPointerErrors []string

	// validateSectionNumber checks the SectionNumber field for validity.
	if sec.SectionNumber == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute SectionNumber cannot be nil")
	} else if *sec.SectionNumber == "" {
		validationErrors = append(validationErrors, "error: attribute SectionNumber cannot be empty")
	}

	// validateCurrentTemperature checks the CurrentTemperature field for validity.
	if sec.CurrentTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentTemperature cannot be nil")
	}

	// validateMinimumTemperature checks the MinimumTemperature field for validity.
	if sec.MinimumTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumTemperature cannot be nil")
	}

	// validateCurrentCapacity checks the CurrentCapacity field for validity.
	if sec.CurrentCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentCapacity cannot be nil")
	} else if *sec.CurrentCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute CurrentCapacity cannot be negative")
	}

	// validateMinimumCapacity checks the MinimumCapacity field for validity.
	if sec.MinimumCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumCapacity cannot be nil")
	} else if *sec.MinimumCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute MinimumCapacity cannot be negative")
	}

	// validateMaximumCapacity checks the MaximumCapacity field for validity.
	if sec.MaximumCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MaximumCapacity cannot be nil")
	} else if *sec.MaximumCapacity <= 0 {
		validationErrors = append(validationErrors, "error: attribute MaximumCapacity must be positive")
	}

	// validateWarehouseID checks the WarehouseID field for validity.
	if sec.WarehouseID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute WarehouseID cannot be nil")
	} else if *sec.WarehouseID <= 0 {
		validationErrors = append(validationErrors, "error: attribute WarehouseID must be positive")
	}

	// validateProductTypeID checks the ProductTypeID field for validity.
	if sec.ProductTypeID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ProductTypeID cannot be nil")
	} else if *sec.ProductTypeID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductTypeID must be positive")
	}

	// Aggregate accumulated errors
	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		allErrors := append(nilPointerErrors, validationErrors...)
		return fmt.Errorf("validation errors: %v", allErrors)
	}

	return nil
}

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
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}

	if err := secReqJSON.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusUnprocessableEntity, err.Error()))

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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok &&
			(mysqlErr.Number == mysqlerr.CodeDuplicateEntry ||
				mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow) {
			response.Error(w, http.StatusConflict, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusConflict, err.Error()))

			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

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
