package sellers_repository

import (
	"strings"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SellersDefault) Update(id int, seller models.SellerDTO) (sellerReturn models.Seller, err error) {
	// Creates update query with the fields that are not empty
	fields := []string{}
	values := []interface{}{}

	if seller.CID != "" {
		fields = append(fields, "cid = ?")
		values = append(values, seller.CID)
	}
	if seller.CompanyName != "" {
		fields = append(fields, "company_name = ?")
		values = append(values, seller.CompanyName)
	}
	if seller.Address != "" {
		fields = append(fields, "address = ?")
		values = append(values, seller.Address)
	}
	if seller.Telephone != "" {
		fields = append(fields, "telephone = ?")
		values = append(values, seller.Telephone)
	}
	if seller.LocalityID != 0 {
		fields = append(fields, "locality_id = ?")
		values = append(values, seller.LocalityID)
	}
	if len(fields) == 0 {
		return
	}

	query := "UPDATE sellers SET " + strings.Join(fields, ", ") + " WHERE id = ?"
	values = append(values, id)

	// Update the seller
	res, err := r.DB.Exec(query, values...)
	// Check for errors
	if err != nil {
		return
	}
	// Check if the seller was updated
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	// If the seller was not updated, return an error
	if rowsAffected == 0 {
		err = models.ErrorIdNotFound
		return
	}

	// Retrieve the updated seller
	err = r.DB.QueryRow("SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id = ?", id).Scan(
		&sellerReturn.ID, &sellerReturn.CID, &sellerReturn.CompanyName, &sellerReturn.Address, &sellerReturn.Telephone, &sellerReturn.LocalityID)
	if err != nil {
		return
	}

	return
}
