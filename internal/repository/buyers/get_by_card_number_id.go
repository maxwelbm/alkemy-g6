package buyersrp

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *BuyerRepository) GetByCardNumberID(cardNumberID string) (buyer models.Buyer, err error) {
	query := "SELECT id, card_number_id, first_name, last_name FROM buyers WHERE card_number_id = ?"

	row := r.db.QueryRow(query, cardNumberID)

	err = row.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrBuyerNotFound
			return
		}
		return
	}

	return
}
