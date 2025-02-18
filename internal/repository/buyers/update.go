package buyersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *BuyerRepository) Update(id int, buyerRequest models.BuyerDTO) (buyer models.Buyer, err error) {
	// Check if the buyer exists
	var exists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM buyers WHERE id = ?)", id).Scan(&exists)

	if err != nil {
		return buyer, err
	}

	if !exists {
		err = models.ErrBuyerNotFound
		return buyer, err
	}

	// Update the buyer
	query := `
	UPDATE buyers SET 
		card_number_id = COALESCE(NULLIF(?, ''), card_number_id), 
		first_name = COALESCE(NULLIF(?, ''), first_name),
		last_name = COALESCE(NULLIF(?, ''), last_name)
	WHERE id = ?`
	_, err = r.db.Exec(query, buyerRequest.CardNumberID, buyerRequest.FirstName, buyerRequest.LastName, id)
	// Check for errors
	if err != nil {
		return buyer, err
	}

	// Retrieve the updated buyer
	err = r.db.QueryRow("SELECT id, card_number_id, first_name, last_name FROM buyers WHERE id = ?", id).Scan(
		&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)

	if err != nil {
		return buyer, err
	}

	return buyer, nil
}
