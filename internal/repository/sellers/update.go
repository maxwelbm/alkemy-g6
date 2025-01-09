package sellers_repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SellersDefault) Update(id int, seller models.SellerDTO) (sellerReturn models.Seller, err error) {
	// Update the seller
	query := `UPDATE sellers SET 
		cid = COALESCE(NULLIF(?, ''), cid), 
		company_name = COALESCE(NULLIF(?, ''), company_name),
		address = COALESCE(NULLIF(?, ''), address),
		telephone = COALESCE(NULLIF(?, ''), telephone),
		locality_id = COALESCE(NULLIF(?, 0), locality_id)
	WHERE id = ?`
	res, err := r.DB.Exec(query, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityID, id)
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
		err = models.ErrorNoChangesMade
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
