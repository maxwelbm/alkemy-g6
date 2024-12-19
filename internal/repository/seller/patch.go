package repository

import modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"

func (r *SellerRepository) PatchSeller(seller modelsSeller.Seller) error {
	r.Sellers[seller.ID] = seller
	return nil
}
