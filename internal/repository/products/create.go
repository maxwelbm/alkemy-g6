package repository

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

func (p *Products) Create(prod models.ProductDTO) (newProd models.Product, err error) {
	query := "INSERT INTO products (`product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := p.DB.Exec(query,
		prod.ProductCode,
		prod.Description,
		prod.Height,
		prod.Length,
		prod.Width,
		prod.NetWeight,
		prod.ExpirationRate,
		prod.FreezingRate,
		prod.RecomFreezTemp,
		prod.ProductTypeID,
		prod.SellerID,
	)

	if err != nil {
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return
	}
	query = "SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products WHERE `id` = ?"
	err = p.DB.QueryRow(query, lastInsertID).Scan(
		&newProd.ID,
		&newProd.ProductCode,
		&newProd.Description,
		&newProd.Height,
		&newProd.Length,
		&newProd.Width,
		&newProd.NetWeight,
		&newProd.ExpirationRate,
		&newProd.FreezingRate,
		&newProd.RecomFreezTemp,
		&newProd.ProductTypeID,
		&newProd.SellerID,
	)
	if err != nil {
		return
	}

	return
}
