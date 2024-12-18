package sellerRepository

import (
	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
)

func (r *SellerRepository) GetAll() (sellerMap []modelsSeller.Seller, err error) {

	for _, value := range r.Sellers {
		sellerMap = append(sellerMap, value)
	}

	return
}
