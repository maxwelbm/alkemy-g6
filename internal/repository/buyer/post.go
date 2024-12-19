package repository

import modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"

func (r *BuyerRepository) PostBuyer(buyer modelsBuyer.Buyer) (buyerReturn modelsBuyer.Buyer, err error) {
	buyer.Id = len(r.Buyers) + 1
	r.NextID++
	r.Buyers[buyer.Id] = buyer
	return buyer, nil
}
