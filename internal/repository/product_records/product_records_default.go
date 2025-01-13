package product_records_repository

import (
	"database/sql"
)

type ProductRecordsDefault struct {
	Db *sql.DB
}

func NewProductRecordsRepository(db *sql.DB) *ProductRecordsDefault {
	rp := &ProductRecordsDefault{
		Db: db,
	}
	return rp
}
