package productbatchesrp

import "database/sql"

type ProductBatchesRepository struct {
	db *sql.DB
}

func NewProductBatchesRepository(db *sql.DB) *ProductBatchesRepository {
	return &ProductBatchesRepository{db: db}
}
