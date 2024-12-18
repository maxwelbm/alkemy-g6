package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/products"

func (p *Products) GetById(id int) (prod models.Product, err error) {
	for _, v := range p.db {
		if v.ID == id {
			prod = v
			return
		}
	}
	return
}
