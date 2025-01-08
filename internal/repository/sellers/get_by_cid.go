package sellers_repository

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SellersDefault) GetByCid(cid int) (seller models.Seller, err error) {
	// query to get seller by cid
	query := "SELECT id, cid, company_name, address, telephone FROM sellers WHERE cid = ?"
	row := r.DB.QueryRow(query, cid)

	// scan row into seller
	err = row.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
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
