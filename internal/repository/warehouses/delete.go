package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/warehouses"

func (r *Warehouses) Delete(id int) (err error) {
	_, ok := r.db[id]
	if !ok {
		err = models.ErrWarehouseRepositoryNotFound
	}
	delete(r.db, id)

	return
}