package repository

import (
	"database/sql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type Products struct {
	DB *sql.DB
}

func NewProducts(db *sql.DB) *Products {
	repo := &Products{
		DB: db,
	}
	return repo
}

func (p *Products) validateProduct(prod models.ProductDTO) (err error) {
	var count int

	// validate ProductCode uniqueness
	query := "SELECT COUNT(*) FROM products WHERE `product_code` = ?"
	err = p.DB.QueryRow(query, prod.ProductCode).Scan(&count)
	if err != nil {
		return
	}

	if count > 0 {
		return models.ErrProductUniqueness
	}

	return nil
}
