package repository

import (
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
)

var (
	ErrProductNotFound = errors.New("Product not found")
)

type Products struct {
	db map[int]models.Product
}

func NewProducts(db map[int]models.Product) *Products {
	defaultDb := make(map[int]models.Product)
	if db != nil {
		defaultDb = db
	}
	return &Products{db: defaultDb}
}
