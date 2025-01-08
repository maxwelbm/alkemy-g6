package buyers_repository

import (
	"strings"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *BuyerRepository) Update(buyerRequest models.BuyerDTO) (buyer models.Buyer, err error) {
	fields := []string{}
	values := []interface{}{}

	if buyerRequest.CardNumberId != nil {
		fields = append(fields, "card_number_id = ?")
		values = append(values, *buyerRequest.CardNumberId)
		buyer.CardNumberId = *buyerRequest.CardNumberId
	}
	if buyerRequest.FirstName != nil {
		fields = append(fields, "first_name = ?")
		values = append(values, *buyerRequest.FirstName)
		buyer.FirstName = *buyerRequest.FirstName
	}
	if buyerRequest.LastName != nil {
		fields = append(fields, "last_name = ?")
		values = append(values, *buyerRequest.LastName)
		buyer.LastName = *buyerRequest.LastName
	}

	if len(fields) == 0 {
		return
	}

	query := "UPDATE buyers SET " + strings.Join(fields, ", ") + " WHERE id = ?"
	values = append(values, buyer.Id)

	_, err = r.DB.Exec(query, values...)
	if err != nil {
		return
	}

	return
}
