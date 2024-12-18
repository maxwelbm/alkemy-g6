package repository

func (p *Products) Delete(id int) (err error) {

	_, ok := p.db[id]
	if !ok {
		err = ErrProductNotFound
	}
	delete(p.db, id)
	return
}
