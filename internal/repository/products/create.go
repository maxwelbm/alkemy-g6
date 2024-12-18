package repository

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
)

func (p *Products) Create(prod models.ProductDTO) (newProd models.Product, err error) {
	nextId := p.lastId + 1

	newProd = models.Product{
		ID:             nextId,
		ProductCode:    prod.ProductCode,
		Description:    prod.Description,
		Height:         prod.Height,
		Length:         prod.Length,
		Width:          prod.Width,
		Weight:         prod.Weight,
		ExpirationRate: prod.ExpirationRate,
		FreezingRate:   prod.FreezingRate,
		RecomFreezTemp: prod.RecomFreezTemp,
		ProductTypeID:  prod.ProductTypeID,
		SellerID:       prod.SellerID,
	}

	p.db[nextId] = newProd
	p.lastId = nextId
	return
}
