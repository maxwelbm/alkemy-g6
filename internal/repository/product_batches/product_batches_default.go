package product_batches_repository

import "database/sql"

type ProductBatchesRepository struct {
	db *sql.DB
}

func NewProductBatchesRepository(db *sql.DB) *ProductBatchesRepository {
	return &ProductBatchesRepository{db: db}
}
