package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/products"

func (p *Products) Update(id int, prod models.ProductDTO) (updatedProd models.Product, err error) {
	updatedProd, ok := p.prods[id]
	// Checks if product is present
	if !ok {
		err = ErrProductNotFound
		return
	}

	// Updates attributes that are present
	if prod.ProductCode != "" {
		updatedProd.ProductCode = prod.ProductCode
	}
	if prod.Description != "" {
		updatedProd.Description = prod.Description
	}
	if prod.Height != 0 {
		updatedProd.Height = prod.Height
	}
	if prod.Length != 0 {
		updatedProd.Length = prod.Length
	}
	if prod.Width != 0 {
		updatedProd.Width = prod.Width
	}
	if prod.Weight != 0 {
		updatedProd.Weight = prod.Weight
	}
	if prod.ExpirationRate != 0 {
		updatedProd.ExpirationRate = prod.ExpirationRate
	}
	if prod.FreezingRate != 0 {
		updatedProd.FreezingRate = prod.FreezingRate
	}
	if prod.RecomFreezTemp != 0 {
		updatedProd.RecomFreezTemp = prod.RecomFreezTemp
	}
	if prod.ProductTypeID != 0 {
		updatedProd.ProductTypeID = prod.ProductTypeID
	}
	if prod.SellerID != 0 {
		updatedProd.SellerID = prod.SellerID
	}

	// Validate all attributes at once
	if err = p.validateProduct(updatedProd); err != nil {
		return
	}

	// updated record in the database
	p.prods[id] = updatedProd
	return
}
