package sellerRepository

import modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"

func (r *SellerRepository) FindAll() (sellerMap map[int]modelsSeller.Seller, err error) {
	sellerMap = make(map[int]modelsSeller.Seller)

	for key, value := range r.Sellers {
		sellerMap[key] = value
	}

	return
}
