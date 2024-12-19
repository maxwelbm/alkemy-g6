package repository

func (e *Employees) Delete(id int) (err error) {

	_, ok := e.db[id]
	if !ok {
		err = ErrEmployeesRepositoryNotFound
		return
	}

	delete(e.db, id)
	return
}
