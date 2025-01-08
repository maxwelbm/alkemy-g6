package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (p *Products) GetAll() (list []models.Product, err error) {
	query := "SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `weight`, `expiration_rate`, `freezing_rate`, `recom_freez_temp`, `product_type_id`, `seller_id` FROM products"
	rows, err := p.DB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(
			&product.ID,
			&product.ProductCode,
			&product.Description,
			&product.Height,
			&product.Length,
			&product.Width,
			&product.Weight,
			&product.ExpirationRate,
			&product.FreezingRate,
			&product.RecomFreezTemp,
			&product.ProductTypeID,
			&product.SellerID,
		)
		if err != nil {
			return
		}
		list = append(list, product)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}
