package buyers_repository

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *BuyerRepository) Delete(id int) (err error) {
	query := "DELETE FROM buyers WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		err = models.ErrBuyerNotFound
		return
	}

	return nil
}
