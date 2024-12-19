package repository

import (
	"errors"
	"fmt"

	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
)

var (
	ErrProductNotFound   = errors.New("Product not found")
	ErrProductUniqueness = errors.New("Product attribute uniqueness")
)

type Products struct {
	prods  map[int]models.Product
	lastId int
}

func NewProducts(prods map[int]models.Product) *Products {
	// initializes db map
	defaultDb := make(map[int]models.Product)
	if prods != nil {
		defaultDb = prods
	}

	// assigns last id
	lastId := 0
	for _, p := range prods {
		if lastId < p.ID {
			lastId = p.ID
		}
	}

	return &Products{prods: defaultDb, lastId: lastId}
}

func (p *Products) validateProduct(prod models.Product) (err error) {
	// Uniqueness
	for _, dbProd := range p.prods {
		// skips self
		if prod.ID == dbProd.ID {
			continue
		}

		// validate ProductCode uniqueness
		if prod.ProductCode == dbProd.ProductCode {
			err = errors.Join(err, fmt.Errorf("%w: %s", ErrProductUniqueness, "Product Code"))
		}
	}
	return
}
