package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/products"

func (p *Products) GetById(id int) (prod models.Product, err error) {

	prod, ok := p.db[id]
	if !ok {
		err = ErrProductNotFound
	}
	return
}
