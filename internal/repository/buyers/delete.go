package buyers_repository

import (
	"database/sql"
)

func (r *BuyerRepository) Delete(id int) (err error) {
	query := "DELETE FROM buyers WHERE id = ?"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
