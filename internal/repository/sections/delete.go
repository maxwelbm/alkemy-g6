package repository

func (r *Sections) Delete(id int) (err error) {
	_, ok := r.db[id]

	if !ok {
		err = ErrSectionNotFound
		return
	}

	delete(r.db, id)

	return
}
