package productbatchesctl

import (
	"errors"
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

	if c.BatchNumber == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute BatchNumber cannot be nil")
	} else if *c.BatchNumber == "" {
		validationErrors = append(validationErrors, "error: attribute BatchNumber cannot be empty")
	}
	if c.InitialQuantity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute InitialQuantity cannot be nil")
	} else if *c.InitialQuantity <= 0 {
		validationErrors = append(validationErrors, "error: attribute InitialQuantity cannot be negative")
	}
	if c.CurrentQuantity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentQuantity cannot be nil")
	} else if *c.CurrentQuantity <= 0 {
		validationErrors = append(validationErrors, "error: attribute CurrentQuantity cannot be negative")
	}
	if c.CurrentTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentTemperature cannot be nil")
	}
	if c.MinimumTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumTemperature cannot be nil")
	}
	if c.DueDate == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute DueDate cannot be nil")
	}
	if c.ManufacturingDate == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ManufacturingDate cannot be nil")
	}
	if c.ManufacturingHour == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ManufacturingHour cannot be nil")
	}
	if c.ProductID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ProductID cannot be nil")
	} else if *c.ProductID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductID must be positive")
	}
	if c.SectionID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute SectionID cannot be nil")
	} else if *c.SectionID <= 0 {
		validationErrors = append(validationErrors, "error: attribute SectionID must be positive")
	}

	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, nilPointerErrors...)
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}
	return
}
