package sectionsctl

import (
	"errors"
	"fmt"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type SectionFullJSON struct {
	ID                 int     `json:"id"`
	SectionNumber      string  `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

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

type SectionResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ProductReportResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ReportProductFullJSON struct {
	SectionID     int    `json:"section_id"`
	SectionNumber string `json:"section_number"`
	ProductsCount int    `json:"products_count"`
}

//nolint:cyclomatic
func (sec *NewSectionReqJSON) validateCreate() (err error) {
	var validationErrors, nilPointerErrors []string

	if sec.SectionNumber == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute SectionNumber cannot be nil")
	} else if *sec.SectionNumber == "" {
		validationErrors = append(validationErrors, "error: attribute SectionNumber cannot be empty")
	}

	if sec.CurrentTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentTemperature cannot be nil")
	}

	if sec.MinimumTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumTemperature cannot be nil")
	}

	if sec.CurrentCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentCapacity cannot be nil")
	} else if *sec.CurrentCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute CurrentCapacity cannot be negative")
	}

	if sec.MinimumCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumCapacity cannot be nil")
	} else if *sec.MinimumCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute MinimumCapacity cannot be negative")
	}

	if sec.MaximumCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MaximumCapacity cannot be nil")
	} else if *sec.MaximumCapacity <= 0 {
		validationErrors = append(validationErrors, "error: attribute MaximumCapacity cannot be negative")
	}

	if sec.WarehouseID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute WarehouseID cannot be nil")
	} else if *sec.WarehouseID <= 0 {
		validationErrors = append(validationErrors, "error: attribute WarehouseID must be positive")
	}

	if sec.ProductTypeID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ProductTypeID cannot be nil")
	} else if *sec.ProductTypeID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductTypeID must be positive")
	}

	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, nilPointerErrors...)
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}

	return err
}

func (sec *NewSectionReqJSON) validateUpdate() (err error) {
	var validationErrors []string

	// Check for nil pointers and collect their errors
	if sec.SectionNumber != nil && *sec.SectionNumber == "" {
		validationErrors = append(validationErrors, "error: attribute SectionNumber cannot be empty")
	}

	if sec.CurrentCapacity != nil && *sec.CurrentCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute CurrentCapacity cannot be negative")
	}

	if sec.MinimumCapacity != nil && *sec.MinimumCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute MinimumCapacity cannot be negative")
	}

	if sec.MaximumCapacity != nil && *sec.MaximumCapacity <= 0 {
		validationErrors = append(validationErrors, "error: attribute MaximumCapacity cannot be negative")
	}

	if sec.WarehouseID != nil && *sec.WarehouseID <= 0 {
		validationErrors = append(validationErrors, "error: attribute WarehouseID must be positive")
	}

	if sec.ProductTypeID != nil && *sec.ProductTypeID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductTypeID must be positive")
	}

	// Combine all errors before returning
	if len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}

	return err
}

type SectionsController struct {
	sv models.SectionService
}

func NewSectionsController(sv models.SectionService) *SectionsController {
	return &SectionsController{sv: sv}
}
