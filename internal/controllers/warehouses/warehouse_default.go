package warehousesctl

import (
	"fmt"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type WarehouseDefault struct {
	sv models.WarehouseService
}

func NewWarehousesController(sv models.WarehouseService) *WarehouseDefault {
	return &WarehouseDefault{sv: sv}
}

type WarehouseDataResJSON struct {
	ID                 int     `json:"id"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
}

type WarehouseResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type WarehouseReqJSON struct {
	Address            *string  `json:"address"`
	Telephone          *string  `json:"telephone"`
	WarehouseCode      *string  `json:"warehouse_code"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
}

func validateNewWarehouse(warehouse WarehouseReqJSON) (err error) {
	var validationErrors, nilPointerErrors []string

	if warehouse.Address == nil {
		nilPointerErrors = append(nilPointerErrors, "error: the address field cannot be nil")
	} else if *warehouse.Address == "" {
		validationErrors = append(validationErrors, "error: the address field cannot be empty")
	}

	if warehouse.Telephone == nil {
		nilPointerErrors = append(nilPointerErrors, "error: the telephone field cannot be nil")
	} else if *warehouse.Telephone == "" {
		validationErrors = append(validationErrors, "error: the telephone field cannot be empty")
	}

	if warehouse.WarehouseCode == nil {
		nilPointerErrors = append(nilPointerErrors, "error: the warehouse_code field cannot be nil")
	} else if *warehouse.WarehouseCode == "" {
		validationErrors = append(validationErrors, "error: the warehouse_code field cannot be empty")
	}

	if warehouse.MinimumCapacity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: the minimum_capacity field cannot be nil")
	} else if *warehouse.MinimumCapacity <= 0 {
		validationErrors = append(validationErrors, "error: the minimum_capacity field must be greater than zero")
	}

	if warehouse.MinimumTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: the minimum_temperature field cannot be nil")
	}

	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, nilPointerErrors...)
		allErrors = append(allErrors, validationErrors...)

		err = fmt.Errorf("validation errors: %v", allErrors)
	}

	return err
}

func validateUpdateWarehouse(w WarehouseReqJSON) (err error) {
	var validationErrors []string

	if w.Address != nil && *w.Address == "" {
		validationErrors = append(validationErrors, "error: the address field cannot be empty")
	}

	if w.Telephone != nil && *w.Telephone == "" {
		validationErrors = append(validationErrors, "error: the telephone field cannot be empty")
	}

	if w.WarehouseCode != nil && *w.WarehouseCode == "" {
		validationErrors = append(validationErrors, "error: the warehouse_code field cannot be empty")
	}

	if w.MinimumCapacity != nil && *w.MinimumCapacity <= 0 {
		if *w.MinimumCapacity <= 0 {
			validationErrors = append(validationErrors, "error: the minimum_capacity field must be greater than zero")
		}
	}

	if w.Address == nil && w.Telephone == nil && w.WarehouseCode == nil && w.MinimumCapacity == nil && w.MinimumTemperature == nil {
		validationErrors = append(validationErrors, "error: at least one field should be present")
	}

	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	return
}
