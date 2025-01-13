package product_records_repository

import (
	"database/sql"
)

type ProductRecordsDefault struct {
	db *sql.DB
}

func NewProductRecordsRepository(db *sql.DB) *ProductRecordsDefault {
	rp := &ProductRecordsDefault{
		db: db,
	}
	return rp
}
