package productbatchesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *ProductBatchesRepository) Create(prodBatches models.ProductBatchesDTO) (newProdBatches models.ProductBatches, err error) {
	query := `INSERT INTO product_batches (batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature,
		due_date, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query, prodBatches.BatchNumber, prodBatches.InitialQuantity, prodBatches.CurrentQuantity,
		prodBatches.CurrentTemperature, prodBatches.MinimumTemperature, prodBatches.DueDate, prodBatches.ManufacturingDate,
		prodBatches.ManufacturingHour, prodBatches.ProductID, prodBatches.SectionID)
	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	query = `SELECT id, batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature, due_date,
		manufacturing_date, manufacturing_hour, product_id, section_id FROM product_batches WHERE id = ?`
	err = r.db.QueryRow(query, lastInsertId).
		Scan(&newProdBatches.ID, &newProdBatches.BatchNumber, &newProdBatches.InitialQuantity, &newProdBatches.CurrentQuantity,
			&newProdBatches.CurrentTemperature, &newProdBatches.MinimumTemperature, &newProdBatches.DueDate, &newProdBatches.ManufacturingDate,
			&newProdBatches.ManufacturingHour, &newProdBatches.ProductID, &newProdBatches.SectionID)

	return
}
