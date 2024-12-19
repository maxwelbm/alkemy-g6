package buyerRepository

import (
	"errors"
	"fmt"

	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

func (r *BuyerRepository) GetByCardNumberId(cardNumberId string) (buyer modelsBuyer.Buyer, err error) {
	for _, value := range r.Buyers {
		if value.CardNumberId == cardNumberId {
			valueConverted := modelsBuyer.Buyer{
				Id:           value.Id,
				CardNumberId: value.CardNumberId,
				FirstName:    value.FirstName,
				LastName:     value.LastName,
			}
			return valueConverted, nil
		}
	}
	return modelsBuyer.Buyer{}, errors.New(fmt.Sprintf("Card Number Id %s not found in the base!", cardNumberId))
}
