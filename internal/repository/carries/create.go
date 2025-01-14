package carriesrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *CarriesDefault) Create(carry models.CarryDTO) (carryToReturn models.Carry, err error) {
	//  insert carry into database
	query := "INSERT INTO carries (cid, company_name, address, phone_number, locality_id) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, carry.CID, carry.CompanyName, carry.Address, carry.PhoneNumber, carry.LocalityID)
	if err != nil {
		return
	}

	// get last inserted id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	// get created carry from database
	query = "SELECT id, cid, company_name, address, phone_number, locality_id FROM carries WHERE id = ?"
	err = r.db.
		QueryRow(query, lastInsertId).
		Scan(&carryToReturn.ID, &carryToReturn.CID, &carryToReturn.CompanyName, &carryToReturn.Address, &carryToReturn.PhoneNumber, &carryToReturn.LocalityID)
	if err != nil {
		return
	}

	return
}
