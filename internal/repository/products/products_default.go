package repository

import (
	"errors"

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
