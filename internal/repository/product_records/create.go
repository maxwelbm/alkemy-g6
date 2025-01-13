package product_records_repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *ProductRecordsDefault) Create(productRecord models.ProductRecordDTO) (productRecordToReturn models.ProductRecord, err error) {
	// Insert productRecord into database
	query := "INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)"
	result, err := r.Db.Exec(query, productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)
	if err != nil {
		return
	}

	// Get last inserted id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	// Get created productRecord from database
	query = "SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records WHERE id = ?"
	err = r.Db.
		QueryRow(query, lastInsertId).
		Scan(&productRecordToReturn.ID, &productRecordToReturn.LastUpdateDate, &productRecordToReturn.PurchasePrice, &productRecordToReturn.SalePrice, &productRecordToReturn.ProductId)
	if err != nil {
		return
	}

	return
}
