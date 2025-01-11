package sections_controller

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
	SectionId     int    `json:"section_id"`
	SectionNumber string `json:"section_number"`
	ProductsCount int    `json:"products_count"`
}

func (c *NewSectionReqJSON) validateCreate() (err error) {
	var validationErrors []string
	var nilPointerErrors []string

	// Check for nil pointers and collect their errors
	if c.SectionNumber == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute SectionNumber cannot be nil")
	} else if *c.SectionNumber == "" {
		validationErrors = append(validationErrors, "error: attribute SectionNumber cannot be empty")
	}

	if c.CurrentTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentTemperature cannot be nil")
	}

	if c.MinimumTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumTemperature cannot be nil")
	}

	if c.CurrentCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentCapacity cannot be nil")
	} else if *c.CurrentCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute CurrentCapacity cannot be negative")
	}

	if c.MinimumCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumCapacity cannot be nil")
	} else if *c.MinimumCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute MinimumCapacity cannot be negative")
	}

	if c.MaximumCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MaximumCapacity cannot be nil")
	} else if *c.MaximumCapacity <= 0 {
		validationErrors = append(validationErrors, "error: attribute MaximumCapacity cannot be negative")
	}

	if c.WarehouseID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute WarehouseID cannot be nil")
	} else if *c.WarehouseID <= 0 {
		validationErrors = append(validationErrors, "error: attribute WarehouseID must be positive")
	}

	if c.ProductTypeID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ProductTypeID cannot be nil")
	} else if *c.ProductTypeID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductTypeID must be positive")
	}

	// Combine all errors before returning
	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, nilPointerErrors...)
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}
	return
}

func (c *NewSectionReqJSON) validateUpdate() (err error) {
	var validationErrors []string

	// Check for nil pointers and collect their errors
	if c.SectionNumber != nil && *c.SectionNumber == "" {
		validationErrors = append(validationErrors, "error: attribute SectionNumber cannot be empty")
	}

	if c.CurrentCapacity != nil && *c.CurrentCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute CurrentCapacity cannot be negative")
	}

	if c.MinimumCapacity != nil && *c.MinimumCapacity < 0 {
		validationErrors = append(validationErrors, "error: attribute MinimumCapacity cannot be negative")
	}

	if c.MaximumCapacity != nil && *c.MaximumCapacity <= 0 {
		validationErrors = append(validationErrors, "error: attribute MaximumCapacity cannot be negative")
	}

	if c.WarehouseID != nil && *c.WarehouseID <= 0 {
		validationErrors = append(validationErrors, "error: attribute WarehouseID must be positive")
	}

	if c.ProductTypeID != nil && *c.ProductTypeID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductTypeID must be positive")
	}

	// Combine all errors before returning
	if len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}
	return
}

type SectionsController struct {
	SV models.SectionService
}

func NewSectionsController(SV models.SectionService) *SectionsController {
	return &SectionsController{SV: SV}
}
