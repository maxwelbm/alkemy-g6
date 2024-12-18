package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/products"

func (p *Products) GetAll() (list []models.Product, err error) {
	for _, v := range p.db {
		list = append(list, v)
	}

	return
}
