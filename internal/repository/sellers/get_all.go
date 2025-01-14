package sellersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SellersDefault) GetAll() (sellers []models.Seller, err error) {
	// query to get all sellers
	query := "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers"
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		var seller models.Seller
		if err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID); err != nil {
			return
		}
		sellers = append(sellers, seller)
	}

	// Check for errors after rows iteration
	if err = rows.Err(); err != nil {
		return
	}

	return
}
