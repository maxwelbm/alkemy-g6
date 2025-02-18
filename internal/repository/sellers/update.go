package sellersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SellersDefault) Update(id int, seller models.SellerDTO) (sellerReturn models.Seller, err error) {
	// Check if the seller exists
	var exists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sellers WHERE id = ?)", id).Scan(&exists)

	if err != nil {
		return sellerReturn, err
	}

	if !exists {
		err = models.ErrSellerNotFound
		return sellerReturn, err
	}
	// Update the seller
	query := `UPDATE sellers SET 
		cid = COALESCE(NULLIF(?, ''), cid), 
		company_name = COALESCE(NULLIF(?, ''), company_name),
		address = COALESCE(NULLIF(?, ''), address),
		telephone = COALESCE(NULLIF(?, ''), telephone),
		locality_id = COALESCE(NULLIF(?, 0), locality_id)
	WHERE id = ?`
	res, err := r.db.Exec(query, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityID, id)
	// Check for errors
	if err != nil {
		return sellerReturn, err
	}
	// Check if the seller was updated
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return sellerReturn, err
	}
	// If the seller was not updated, return an error
	if rowsAffected == 0 {
		err = models.ErrorNoChangesMade
		return sellerReturn, err
	}

	// Retrieve the updated seller
	err = r.db.QueryRow("SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id = ?", id).Scan(
		&sellerReturn.ID, &sellerReturn.CID, &sellerReturn.CompanyName, &sellerReturn.Address, &sellerReturn.Telephone, &sellerReturn.LocalityID)
	if err != nil {
		return sellerReturn, err
	}

	return sellerReturn, err
}
