package sellersrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *SellersDefault) Create(seller models.SellerDTO) (sellerToReturn models.Seller, err error) {
	//  insert seller into database
	query := "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityID)
	if err != nil {
		return
	}

	// get last inserted id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	// get created seller from database
	query = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id = ?"
	err = r.db.
		QueryRow(query, lastInsertId).
		Scan(&sellerToReturn.ID, &sellerToReturn.CID, &sellerToReturn.CompanyName, &sellerToReturn.Address, &sellerToReturn.Telephone, &sellerToReturn.LocalityID)
	if err != nil {
		return
	}

	return
}
