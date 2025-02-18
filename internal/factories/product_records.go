package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type ProductRecordsFactory struct {
	db *sql.DB
}

func NewProductRecordsFactory(db *sql.DB) *ProductRecordsFactory {
	return &ProductRecordsFactory{db: db}
}

func defaultProductRecords() models.ProductRecord {
	return models.ProductRecord{
		LastUpdateDate: randstr.Date(),
		PurchasePrice:  10.0,
		SalePrice:      10.0,
		ProductID:      1,
	}
}

func (f *ProductRecordsFactory) Build(productRecord models.ProductRecord) models.ProductRecord {
	populateProductRecordsParams(&productRecord)

	return productRecord
}

func (f *ProductRecordsFactory) Create(productRecord models.ProductRecord) (record models.ProductRecord, err error) {
	populateProductRecordsParams(&productRecord)

	if err = f.checkProductExists(productRecord.ProductID); err != nil {
		return productRecord, err
	}

	query := `
		INSERT INTO product_records 
			(
			%s
			last_update_date,
			purchase_price, 
			sale_price, 
			product_id
			) 
		VALUES (%s?, ?, ?, ?)
	`

	switch productRecord.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(productRecord.ID)+",")
	}

	_, err = f.db.Exec(query,
		productRecord.LastUpdateDate,
		productRecord.PurchasePrice,
		productRecord.SalePrice,
		productRecord.ProductID)

	return productRecord, err
}

func populateProductRecordsParams(productRecord *models.ProductRecord) {
	defaultProductRecords := defaultProductRecords()
	if productRecord == nil {
		productRecord = &defaultProductRecords
	}

	if productRecord.LastUpdateDate == "" {
		productRecord.LastUpdateDate = defaultProductRecords.LastUpdateDate
	}

	if productRecord.PurchasePrice == 0.0 {
		productRecord.PurchasePrice = defaultProductRecords.PurchasePrice
	}

	if productRecord.SalePrice == 0.0 {
		productRecord.SalePrice = defaultProductRecords.SalePrice
	}

	if productRecord.ProductID == 0 {
		productRecord.ProductID = defaultProductRecords.ProductID
	}
}

func (f *ProductRecordsFactory) checkProductExists(productID int) (err error) {
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

func (f *ProductRecordsFactory) createProduct() (err error) {
	productFactory := NewProductFactory(f.db)
	_, err = productFactory.Create(models.Product{})

	return
}
