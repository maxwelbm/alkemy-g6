package repository

import modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"

func (r *SellerRepository) PostSeller(seller modelsSeller.Seller) (sellerToReturn modelsSeller.Seller, err error) {
	nextId := r.LastId + 1
	seller.ID = nextId
	r.Sellers[seller.ID] = seller
	sellerToReturn = seller
	return
}
