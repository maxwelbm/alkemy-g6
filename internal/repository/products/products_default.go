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
	db     map[int]models.Product
	lastId int
}

func NewProducts(db map[int]models.Product) *Products {
	// initializes db map
	defaultDb := make(map[int]models.Product)
	if db != nil {
		defaultDb = db
	}

	// assigns last id
	lastId := 0
	for _, p := range db {
		if lastId < p.ID {
			lastId = p.ID
		}
	}

	return &Products{db: defaultDb, lastId: lastId}
}

func (p *Products) validateProduct(prod models.Product) (err error) {
	// Uniqueness
	for _, dbProd := range p.db {
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
