package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type ProductBatchesFactory struct {
	db *sql.DB
}

func NewProductBatchesFactory(db *sql.DB) *ProductBatchesFactory {
	return &ProductBatchesFactory{db: db}
}

func defaultProductBatches() models.ProductBatches {
	return models.ProductBatches{
		BatchNumber:        randstr.Alphanumeric(8),
		InitialQuantity:    1,
		CurrentQuantity:    1,
		CurrentTemperature: 10.0,
		MinimumTemperature: 1.0,
		DueDate:            randstr.Data(),
		ManufacturingDate:  randstr.Data(),
		ManufacturingHour:  "10:00",
		ProductID:          1,
		SectionID:          1,
	}
}
func (f *ProductBatchesFactory) Build(productBatches models.ProductBatches) models.ProductBatches {
	populateProductBatchesParams(&productBatches)

	return productBatches
}

func (f *ProductBatchesFactory) Create(productBatches models.ProductBatches) (record models.ProductBatches, err error) {
	populateProductBatchesParams(&productBatches)

	if err = f.checkProductExists(productBatches.ProductID); err != nil {
		return productBatches, err
	}

	if err = f.checkSectionExists(productBatches.SectionID); err != nil {
		return productBatches, err
	}

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

func (f *ProductBatchesFactory) checkProductExists(productID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM products WHERE id = ?`, productID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createProduct()

	return
}

func (f *ProductBatchesFactory) checkSectionExists(sectionID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM sections WHERE id = ?`, sectionID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createSections()

	return
}

func (f *ProductBatchesFactory) createProduct() (err error) {
	productBatchesFactory := NewProductFactory(f.db)
	_, err = productBatchesFactory.Create(models.Product{})

	return
}

func (f *ProductBatchesFactory) createSections() (err error) {
	productBatchesFactory := NewSectionFactory(f.db)
	_, err = productBatchesFactory.Create(models.Section{})

	return
}
