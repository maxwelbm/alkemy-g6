package sellersrp

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SellersDefault) GetByCid(cid int) (seller models.Seller, err error) {
	// query to get seller by cid
	query := "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE cid = ?"
	row := r.db.QueryRow(query, cid)

	// scan row into seller
	err = row.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityID)
	// check for errors
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrSellerNotFound
			return
		}

		return
	}

	return
}
