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

type SectionsController struct {
	sv models.SectionService
}

func NewSectionsController(sv models.SectionService) *SectionsController {
	return &SectionsController{sv: sv}
}

//nolint:cyclomatic
func (sec *NewSectionReqJSON) validateCreate() (err error) {
	var validationErrors []string

	var nilPointerErrors []string

	if err := sec.validateSectionNumber(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := sec.validateCurrentTemperature(&nilPointerErrors); err != nil {
		return err
	}

	if err := sec.validateMinimumTemperature(&nilPointerErrors); err != nil {
		return err
	}

	if err := sec.validateCurrentCapacity(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := sec.validateMinimumCapacity(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := sec.validateMaximumCapacity(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := sec.validateWarehouseID(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := sec.validateProductTypeID(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	// Aggregate accumulated errors
	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		allErrors := append(nilPointerErrors, validationErrors...)
		return fmt.Errorf("validation errors: %v", allErrors)
	}

	return nil
}

// validateSectionNumber checks the SectionNumber field for validity.
func (sec *NewSectionReqJSON) validateSectionNumber(validationErrors, nilPointerErrors *[]string) (err error) {
	if sec.SectionNumber == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute SectionNumber cannot be nil")
	} else if *sec.SectionNumber == "" {
		*validationErrors = append(*validationErrors, "error: attribute SectionNumber cannot be empty")
	}

	return
}

// validateCurrentTemperature checks the CurrentTemperature field for validity.
func (sec *NewSectionReqJSON) validateCurrentTemperature(nilPointerErrors *[]string) (err error) {
	if sec.CurrentTemperature == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute CurrentTemperature cannot be nil")
	}

	return
}

// validateMinimumTemperature checks the MinimumTemperature field for validity.
func (sec *NewSectionReqJSON) validateMinimumTemperature(nilPointerErrors *[]string) (err error) {
	if sec.MinimumTemperature == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute MinimumTemperature cannot be nil")
	}

	return
}

// validateCurrentCapacity checks the CurrentCapacity field for validity.
func (sec *NewSectionReqJSON) validateCurrentCapacity(validationErrors, nilPointerErrors *[]string) (err error) {
	if sec.CurrentCapacity == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute CurrentCapacity cannot be nil")
	} else if *sec.CurrentCapacity < 0 {
		*validationErrors = append(*validationErrors, "error: attribute CurrentCapacity cannot be negative")
	}

	return
}

// validateMinimumCapacity checks the MinimumCapacity field for validity.
func (sec *NewSectionReqJSON) validateMinimumCapacity(validationErrors, nilPointerErrors *[]string) (err error) {
	if sec.MinimumCapacity == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute MinimumCapacity cannot be nil")
	} else if *sec.MinimumCapacity < 0 {
		*validationErrors = append(*validationErrors, "error: attribute MinimumCapacity cannot be negative")
	}

	return
}

// validateMaximumCapacity checks the MaximumCapacity field for validity.
func (sec *NewSectionReqJSON) validateMaximumCapacity(validationErrors, nilPointerErrors *[]string) (err error) {
	if sec.MaximumCapacity == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute MaximumCapacity cannot be nil")
	} else if *sec.MaximumCapacity <= 0 {
		*validationErrors = append(*validationErrors, "error: attribute MaximumCapacity must be positive")
	}

	return
}

// validateWarehouseID checks the WarehouseID field for validity.
func (sec *NewSectionReqJSON) validateWarehouseID(validationErrors, nilPointerErrors *[]string) (err error) {
	if sec.WarehouseID == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute WarehouseID cannot be nil")
	} else if *sec.WarehouseID <= 0 {
		*validationErrors = append(*validationErrors, "error: attribute WarehouseID must be positive")
	}

	return
}

// validateProductTypeID checks the ProductTypeID field for validity.
func (sec *NewSectionReqJSON) validateProductTypeID(validationErrors, nilPointerErrors *[]string) (err error) {
	if sec.ProductTypeID == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute ProductTypeID cannot be nil")
	} else if *sec.ProductTypeID <= 0 {
		*validationErrors = append(*validationErrors, "error: attribute ProductTypeID must be positive")
	}

	return
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
