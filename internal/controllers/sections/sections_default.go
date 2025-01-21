package sectionsctl

import (
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

type SectionsController struct {
	sv models.SectionService
}

func NewSectionsController(sv models.SectionService) *SectionsController {
	return &SectionsController{sv: sv}
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

		err = fmt.Errorf("validation errors: %v", allErrors)
	}

	return err
}
