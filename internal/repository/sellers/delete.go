package sellersrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *SellersDefault) Delete(id int) (err error) {
	query := "DELETE FROM sellers WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		err = models.ErrSellerNotFound
		return
	}

	return nil
}
