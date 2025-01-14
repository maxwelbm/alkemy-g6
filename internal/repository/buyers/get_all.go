package buyersrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *BuyerRepository) GetAll() (buyers []models.Buyer, err error) {
	query := "SELECT id, card_number_id, first_name, last_name FROM buyers"

	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var buyer models.Buyer
		if err = rows.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName); err != nil {
			return
		}
		buyers = append(buyers, buyer)
	}

	// Check for errors after rows iteration
	if err = rows.Err(); err != nil {
		return
	}

	return
}
