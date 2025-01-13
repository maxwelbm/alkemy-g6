package warehouses_repository

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *WarehouseRepository) Delete(id int) (err error) {
	query := "DELETE FROM warehouses WHERE `id`=?"
	result, err := r.db.Exec(query, id)

	if err != nil {
		return
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowAffected == 0 {
		err = models.ErrWareHouseNotFound
		return
	}

	return
}
