package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"strings"
)

func (p *Products) Update(id int, prod models.ProductDTO) (updatedProd models.Product, err error) {
	// Validate all attributes
	if err = p.validateProduct(prod); err != nil {
		return
	}

	// Prepare fields and values for the update query
	fields := []string{}
	values := []interface{}{}

	if prod.ProductCode != "" {
		fields = append(fields, "product_code = ?")
		values = append(values, prod.ProductCode)
	}
	if prod.Description != "" {
		fields = append(fields, "description = ?")
		values = append(values, prod.Description)
	}
	if prod.Height != 0 {
		fields = append(fields, "height = ?")
		values = append(values, prod.Height)
	}
	if prod.Length != 0 {
		fields = append(fields, "length = ?")
		values = append(values, prod.Length)
	}
	if prod.Width != 0 {
		fields = append(fields, "width = ?")
		values = append(values, prod.Width)
	}
	if prod.Weight != 0 {
		fields = append(fields, "weight = ?")
		values = append(values, prod.Weight)
	}
	if prod.ExpirationRate != 0 {
		fields = append(fields, "expiration_rate = ?")
		values = append(values, prod.ExpirationRate)
	}
	if prod.FreezingRate != 0 {
		fields = append(fields, "freezing_rate = ?")
		values = append(values, prod.FreezingRate)
	}
	if prod.RecomFreezTemp != 0 {
		fields = append(fields, "recom_freez_temp = ?")
		values = append(values, prod.RecomFreezTemp)
	}
	if prod.ProductTypeID != 0 {
		fields = append(fields, "product_type_id = ?")
		values = append(values, prod.ProductTypeID)
	}
	if prod.SellerID != 0 {
		fields = append(fields, "seller_id = ?")
		values = append(values, prod.SellerID)
	}

	// If no fields are specified, nothing to update
	if len(fields) == 0 {
		return 
	}

	query := "UPDATE products SET " + strings.Join(fields, ", ") + " WHERE id = ?"
	values = append(values, id)

	// Execute the update query
	res, err := p.DB.Exec(query, values...)
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

	query = "SELECT id, product_code, description, height, length, width, weight, expiration_rate, freezing_rate, recom_freez_temp, product_type_id, seller_id FROM products WHERE id = ?"
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
