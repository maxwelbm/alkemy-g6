package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type ProductBatchesFactory struct {
	db *sql.DB
}

func NewProductBatchesFactory(db *sql.DB) *ProductBatchesFactory {
	return &ProductBatchesFactory{db: db}
}

func defaultProductBatches() models.ProductBatches {
	return models.ProductBatches{}
}
func (f *ProductBatchesFactory) Build(productBatches models.ProductBatches) models.ProductBatches {
	populateProductBatchesParams(&productBatches)

	return productBatches
}

func (f *ProductBatchesFactory) Create(productBatches models.ProductBatches) (record models.ProductBatches, err error) {
	populateProductBatchesParams(&productBatches)

	query := `
		INSERT INTO product_batches
			(
			%s
			batch_number,
			initial_quantity,
			current_quantity,
			current_temperature,
			minimum_temperature,
			due_date,
			manufacturing_date,
			manufacturing_hour,
			product_id,
			section_id,
			) 
		VALUES (%s?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	switch productBatches.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(productBatches.ID)+",")
	}

	result, err := f.db.Exec(query,
		productBatches.BatchNumber,
		productBatches.InitialQuantity,
		productBatches.CurrentQuantity,
		productBatches.CurrentTemperature,
		productBatches.MinimumTemperature,
		productBatches.DueDate,
		productBatches.ManufacturingDate,
		productBatches.ManufacturingHour,
		productBatches.ProductID,
		productBatches.SectionID,
	)

	if err != nil {
		return record, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return record, err
	}

	productBatches.ID = int(id)

	return productBatches, err
}

func populateProductBatchesParams(productBatches *models.ProductBatches) {
	defaultProductBatches := defaultProductBatches()
	if productBatches == nil {
		productBatches = &defaultProductBatches
	}

	if productBatches.BatchNumber == "" {
		productBatches.BatchNumber = defaultProductBatches.BatchNumber
	}

	if productBatches.InitialQuantity == 0 {
		productBatches.InitialQuantity = defaultProductBatches.InitialQuantity
	}

	if productBatches.CurrentQuantity == 0 {
		productBatches.CurrentQuantity = defaultProductBatches.CurrentQuantity
	}

	if productBatches.CurrentTemperature == 0 {
		productBatches.CurrentTemperature = defaultProductBatches.CurrentTemperature
	}

	if productBatches.MinimumTemperature == 0 {
		productBatches.MinimumTemperature = defaultProductBatches.MinimumTemperature
	}

	if productBatches.DueDate == "" {
		productBatches.DueDate = defaultProductBatches.DueDate
	}

	if productBatches.ManufacturingDate == "" {
		productBatches.ManufacturingDate = defaultProductBatches.ManufacturingDate
	}

	if productBatches.ManufacturingHour == "" {
		productBatches.ManufacturingHour = defaultProductBatches.ManufacturingHour
	}

	if productBatches.ProductID == 0 {
		productBatches.ProductID = defaultProductBatches.ProductID
	}

	if productBatches.SectionID == 0 {
		productBatches.SectionID = defaultProductBatches.SectionID
	}
}
