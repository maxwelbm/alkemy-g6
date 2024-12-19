package repository

func (p *Products) Delete(id int) (err error) {

	_, ok := p.prods[id]
	if !ok {
		err = ErrProductNotFound
	}
	delete(p.prods, id)
	return
}
