package buyersrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *BuyerRepository) Create(buyer models.BuyerDTO) (buyerReturn models.Buyer, err error) {
	query := "INSERT INTO buyers (card_number_id, first_name, last_name) VALUES (?, ?, ?)"

	result, err := r.db.Exec(query, buyer.CardNumberId, buyer.FirstName, buyer.LastName)
	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	query = "SELECT id, card_number_id, first_name, last_name FROM buyers WHERE id = ?"
	err = r.db.
		QueryRow(query, lastInsertId).
		Scan(&buyerReturn.Id, &buyerReturn.CardNumberId, &buyerReturn.FirstName, &buyerReturn.LastName)
	if err != nil {
		return
	}

	return
}
