package repository

import (
	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

func (r *BuyerRepository) PatchBuyer(buyerRequest modelsBuyer.BuyerDTO) (buyer modelsBuyer.Buyer, err error) {

	buyer, err = r.GetById(*buyerRequest.Id)

	if err != nil {
		err = ErrorIdNotFound
		return
	}

	if buyerRequest.Id != nil {
		buyer.Id = *buyerRequest.Id
	}
	if buyerRequest.CardNumberId != nil {
		buyer.CardNumberId = *buyerRequest.CardNumberId
	}
	if buyerRequest.FirstName != nil {
		buyer.FirstName = *buyerRequest.FirstName
	}
	if buyerRequest.LastName != nil {
		buyer.LastName = *buyerRequest.LastName
	}

	r.Buyers[buyer.Id] = buyer
	return
}
