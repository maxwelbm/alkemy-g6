package buyers_repository

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *BuyerRepository) GetByCardNumberId(cardNumberId string) (buyer models.Buyer, err error) {
	query := "SELECT id, card_number_id, first_name, last_name FROM buyers WHERE card_number_id = ?"

	row := r.DB.QueryRow(query, cardNumberId)

	err = row.Scan(&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrBuyerNotFound
			return
		}
		return
	}

	return
}
