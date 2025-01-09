package buyers_repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *BuyerRepository) Update(id int, buyerRequest models.BuyerDTO) (buyer models.Buyer, err error) {
	// Check if the buyer exists
	var exists bool
	err = r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM buyers WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return
	}
	if !exists {
		err = models.ErrBuyerNotFound
		return
	}

	// Update the buyer
	query := `UPDATE buyers SET 
		card_number_id = COALESCE(NULLIF(?, ''), card_number_id), 
		first_name = COALESCE(NULLIF(?, ''), first_name),
		last_name = COALESCE(NULLIF(?, ''), last_name)
	WHERE id = ?`
	res, err := r.DB.Exec(query, buyerRequest.CardNumberId, buyerRequest.FirstName, buyerRequest.LastName, id)
	// Check for errors
	if err != nil {
		return
	}
	// Check if the buyer was updated
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	// If the buyer was not updated, return an error
	if rowsAffected == 0 {
		err = models.ErrorNoChangesMade
		return
	}

	// Retrieve the updated buyer
	err = r.DB.QueryRow("SELECT id, card_number_id, first_name, last_name FROM buyers WHERE id = ?", id).Scan(
		&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		return
	}

	return
}
