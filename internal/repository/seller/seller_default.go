package repository

import modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"

type SellerRepository struct {
	Sellers map[int]modelsSeller.Seller
	LastId  int
}

func NewSellerRepository(db map[int]modelsSeller.Seller) *SellerRepository {
	// initializes db map
	defaultDb := make(map[int]modelsSeller.Seller)
	if db != nil {
		defaultDb = db
	}

	// assigns last id
	lastId := 0
	for _, seller := range db {
		if lastId < seller.ID {
			lastId = seller.ID
		}
	}

	return &SellerRepository{Sellers: defaultDb, LastId: lastId}
}
