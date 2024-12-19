package buyerRepository

import modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"

func (r *BuyerRepository) PatchBuyer(buyer modelsBuyer.Buyer) (err error) {
	r.Buyers[buyer.Id] = buyer
	return nil
}
