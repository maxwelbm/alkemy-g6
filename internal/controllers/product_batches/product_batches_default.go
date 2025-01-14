package productbatchesctl

import (
	"fmt"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type ProductBatchesController struct {
	sv models.ProductBatchesService
}

func NewProductBatchesController(sv models.ProductBatchesService) *ProductBatchesController {
	return &ProductBatchesController{sv: sv}
}

type NewProductBatchesReqJSON struct {
	BatchNumber        *string  `json:"batch_number"`
	InitialQuantity    *int     `json:"initial_quantity"`
	CurrentQuantity    *int     `json:"current_quantity"`
	CurrentTemperature *float64 `json:"current_temperature"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
	DueDate            *string  `json:"due_date"`
	ManufacturingDate  *string  `json:"manufacturing_date"`
	ManufacturingHour  *string  `json:"manufacturing_hour"`
	ProductID          *int     `json:"product_id"`
	SectionID          *int     `json:"section_id"`
}

type ProductBatchFullJSON struct {
	ID                 int     `json:"id"`
	BatchNumber        string  `json:"batch_number"`
	InitialQuantity    int     `json:"initial_quantity"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	DueDate            string  `json:"due_date"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_hour"`
	ProductID          int     `json:"product_id"`
	SectionID          int     `json:"section_id"`
}

type ProductBatchesResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (c *NewProductBatchesReqJSON) validateCreate() (err error) {
	var validationErrors []string

	var nilPointerErrors []string

	// Validate fields and append errors if any
	if err := c.validateBatchNumber(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateInitialQuantity(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateCurrentQuantity(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateCurrentTemperature(&nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateMinimumTemperature(&nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateDueDate(&nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateManufacturingDate(&nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateManufacturingHour(&nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateProductID(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	if err := c.validateSectionID(&validationErrors, &nilPointerErrors); err != nil {
		return err
	}

	// Check if there are any accumulated errors
	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		allErrors := append(nilPointerErrors, validationErrors...)
		return fmt.Errorf("validation errors: %v", allErrors)
	}

	return err
}

func (c *NewProductBatchesReqJSON) validateBatchNumber(validationErrors, nilPointerErrors *[]string) (err error) {
	if c.BatchNumber == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute BatchNumber cannot be nil")
	} else if *c.BatchNumber == "" {
		*validationErrors = append(*validationErrors, "error: attribute BatchNumber cannot be empty")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateInitialQuantity(validationErrors, nilPointerErrors *[]string) (err error) {
	if c.InitialQuantity == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute InitialQuantity cannot be nil")
	} else if *c.InitialQuantity <= 0 {
		*validationErrors = append(*validationErrors, "error: attribute InitialQuantity must be positive")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateCurrentQuantity(validationErrors, nilPointerErrors *[]string) (err error) {
	if c.CurrentQuantity == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute CurrentQuantity cannot be nil")
	} else if *c.CurrentQuantity <= 0 {
		*validationErrors = append(*validationErrors, "error: attribute CurrentQuantity must be positive")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateCurrentTemperature(nilPointerErrors *[]string) (err error) {
	if c.CurrentTemperature == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute CurrentTemperature cannot be nil")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateMinimumTemperature(nilPointerErrors *[]string) (err error) {
	if c.MinimumTemperature == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute MinimumTemperature cannot be nil")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateDueDate(nilPointerErrors *[]string) (err error) {
	if c.DueDate == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute DueDate cannot be nil")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateManufacturingDate(nilPointerErrors *[]string) (err error) {
	if c.ManufacturingDate == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute ManufacturingDate cannot be nil")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateManufacturingHour(nilPointerErrors *[]string) (err error) {
	if c.ManufacturingHour == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute ManufacturingHour cannot be nil")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateProductID(validationErrors, nilPointerErrors *[]string) (err error) {
	if c.ProductID == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute ProductID cannot be nil")
	} else if *c.ProductID <= 0 {
		*validationErrors = append(*validationErrors, "error: attribute ProductID must be positive")
	}

	return
}

func (c *NewProductBatchesReqJSON) validateSectionID(validationErrors, nilPointerErrors *[]string) (err error) {
	if c.SectionID == nil {
		*nilPointerErrors = append(*nilPointerErrors, "error: attribute SectionID cannot be nil")
	} else if *c.SectionID <= 0 {
		*validationErrors = append(*validationErrors, "error: attribute SectionID must be positive")
	}

	return
}
