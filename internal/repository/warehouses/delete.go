package repository

func (r *Warehouses) Delete(id int) (err error) {
	_, ok := r.db[id]
	if !ok {
		err = ErrWarehouseRepositoryNotFound
	}
	delete(r.db, id)

	return
}