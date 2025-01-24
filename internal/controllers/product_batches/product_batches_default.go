package productbatchesctl

import (
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
