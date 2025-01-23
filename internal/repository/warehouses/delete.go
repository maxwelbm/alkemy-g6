package warehousesrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *WarehouseRepository) Delete(id int) (err error) {
	var exists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM warehouses WHERE id = ?)", id).Scan(&exists)

	if err != nil {
		return
	}

	if !exists {
		err = models.ErrWareHouseNotFound
		return
	}

	query := "DELETE FROM warehouses WHERE id = ?"
	_, err = r.db.Exec(query, id)

	return
}
