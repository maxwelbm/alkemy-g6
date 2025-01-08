package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"errors"
	"database/sql"
)

func (p *Products) GetById(id int) (prod models.Product, err error) {
	query := "SELECT id, product_code, description, height, length, width, weight, expiration_rate, freezing_rate, recom_freez_temp, product_type_id, seller_id FROM products WHERE id = ?"

	row := p.DB.QueryRow(query, id)
	err = row.Scan(
		&prod.ID,
		&prod.ProductCode,
		&prod.Description,
		&prod.Height,
		&prod.Length,
		&prod.Width,
		&prod.Weight,
		&prod.ExpirationRate,
		&prod.FreezingRate,
		&prod.RecomFreezTemp,
		&prod.ProductTypeID,
		&prod.SellerID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrProductNotFound
			return
		}
		return
	}
	return
}
