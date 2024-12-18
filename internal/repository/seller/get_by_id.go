package sellerRepository

import (
	"errors"
	"fmt"

	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
)

func (r *SellerRepository) GetById(id int) (sellerMap modelsSeller.Seller, err error) {
	for _, value := range r.Sellers {
		if value.ID == id {
			valueConverted := modelsSeller.Seller{
				ID:          value.ID,
				CID:         value.CID,
				CompanyName: value.CompanyName,
				Address:     value.Address,
				Telephone:   value.Telephone,
			}
			return valueConverted, nil
		}
	}
	return modelsSeller.Seller{}, errors.New(fmt.Sprintf("Id %d not found in the base!", id))
}
