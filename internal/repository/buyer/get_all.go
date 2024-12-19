package repository

import (
	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

func (r *BuyerRepository) GetAll() (buyers []modelsBuyer.Buyer, err error) {
	for _, value := range r.Buyers {
		buyers = append(buyers, value)
	}
	return
}
