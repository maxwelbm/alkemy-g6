package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/products"

func (p *Products) GetAll() (list map[int]models.Product, err error) {
	list = make(map[int]models.Product)

	for k, v := range p.db {
		list[k] = v
	}

	return
}
