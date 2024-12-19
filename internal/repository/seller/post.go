package repository

import modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"

func (r *SellerRepository) PostSeller(seller modelsSeller.Seller) error {
	seller.ID = len(r.Sellers) + 1
	r.NextID++
	r.Sellers[seller.ID] = seller
	return nil
}
