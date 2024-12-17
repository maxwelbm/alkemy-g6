package product_repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/products"

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
