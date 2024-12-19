package buyerRepository

import (
	"errors"
	"fmt"

	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

func (r *BuyerRepository) GetByCardNumberId(cardNumberId string) (buyer modelsBuyer.Buyer, err error) {
	for _, value := range r.Buyers {
		if value.CardNumberId == cardNumberId {
			return value, nil
		}
	}
	return buyer, errors.New(fmt.Sprintf("Card Number Id %s not found in the base!", cardNumberId))
}
