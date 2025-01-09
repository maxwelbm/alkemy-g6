package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (p *Products) Update(id int, prod models.ProductDTO) (updatedProd models.Product, err error) {
	if prod.ProductCode != "" {
		var exists bool

		query := "SELECT EXISTS(SELECT 1 FROM products WHERE `product_code`=?)"
		err = p.DB.QueryRow(query, prod.ProductCode).Scan(&exists)
		if err != nil {
			return
		}

		if !exists {
			err = models.ErrProductNotFound
			return
		}
    }

	query := `UPDATE products SET 
			product_code = COALESCE(NULLIF(?, ''), product_code), 
			description = COALESCE(NULLIF(?, ''), description),
			height = COALESCE(NULLIF(?, 0), height),
			length = COALESCE(NULLIF(?, 0), length),
			width = COALESCE(NULLIF(?, 0), width),
			net_weight = COALESCE(NULLIF(?, 0), net_weight),
			expiration_rate = COALESCE(NULLIF(?, 0), expiration_rate),
			freezing_rate = COALESCE(NULLIF(?, 0), freezing_rate),
			recommended_freezing_temperature = COALESCE(NULLIF(?, 0), recommended_freezing_temperature),
			product_type_id = COALESCE(NULLIF(?, 0), product_type_id),
			seller_id = COALESCE(NULLIF(?, 0), seller_id)
			WHERE id = ?`

	// Execute the update query
	res, err := p.DB.Exec(query, 
		prod.ProductCode, 
		prod.Description, 
		prod.Height, 
		prod.Length, 
		prod.Width, 
		prod.Weight, 
		prod.ExpirationRate, 
		prod.FreezingRate, 
		prod.RecomFreezTemp, 
		prod.ProductTypeID, 
		prod.SellerID,
		id)

	if err != nil {
		return
	}
	// Check how many rows were affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	// Check if the update affected any rows
	if rowsAffected == 0 {
		err = models.ErrProductNotFound
		return
	}

	query = "SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products WHERE `id` = ?"
	err = p.DB.QueryRow(query, id).Scan(
		&updatedProd.ID,
		&updatedProd.ProductCode,
		&updatedProd.Description,
		&updatedProd.Height,
		&updatedProd.Length,
		&updatedProd.Width,
		&updatedProd.Weight,
		&updatedProd.ExpirationRate,
		&updatedProd.FreezingRate,
		&updatedProd.RecomFreezTemp,
		&updatedProd.ProductTypeID,
		&updatedProd.SellerID,
	)
	if err != nil {
		return
	}

	return
}
