package repository

import (
	"errors"
	"fmt"

	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
)

var (
	ErrProductNotFound = errors.New("Product not found")
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
	var validationErrors []string

	// validate ProductCode uniqueness
	for _, dbProd := range p.db {
		if prod.ProductCode == dbProd.ProductCode {
			validationErrors = append(validationErrors, "error: attribute ProductCode must be unique")
		}
	}

	// joins all validation errors into a new error
	if len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}
	return
}
