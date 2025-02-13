package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type SellerFactory struct {
	db *sql.DB
}

func NewSellerFactory(db *sql.DB) *SellerFactory {
	return &SellerFactory{db: db}
}

func defaultSeller() models.Seller {
	return models.Seller{
		CID:         randstr.Alphanumeric(8),
		CompanyName: randstr.Chars(8),
		Address:     randstr.Chars(8),
		Telephone:   randstr.Chars(11),
		LocalityID:  1,
	}
}

func (f *SellerFactory) Create(seller models.Seller) (record models.Seller, err error) {
	populateSellerParams(&seller)

	if err = f.checkLocalityExists(seller.LocalityID); err != nil {
		return seller, err
	}

	query := `
		INSERT INTO sellers 
			(
			%s
			cid,
			company_name, 
			address, 
			telephone, 
			locality_id
			) 
		VALUES (%s?, ?, ?, ?, ?)
	`

	switch seller.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(seller.ID)+",")
	}

	_, err = f.db.Exec(query,
		seller.CID,
		seller.CompanyName,
		seller.Address,
		seller.Telephone,
		seller.LocalityID,
	)

	return seller, err
}

func populateSellerParams(seller *models.Seller) {
	defaultSeller := defaultSeller()
	if seller == nil {
		seller = &defaultSeller
	}

	if seller.CID == "" {
		seller.CID = defaultSeller.CID
	}

	if seller.CompanyName == "" {
		seller.CompanyName = defaultSeller.CompanyName
	}

	if seller.Address == "" {
		seller.Address = defaultSeller.Address
	}

	if seller.Telephone == "" {
		seller.Telephone = defaultSeller.Telephone
	}

	if seller.LocalityID == 0 {
		seller.LocalityID = defaultSeller.LocalityID
	}
}

func (f *SellerFactory) checkLocalityExists(localityID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM localities WHERE id = ?`, localityID).Scan(&count)

	if err != nil {
		return
	}

	if count == 0 {
		err = fmt.Errorf("seller with id %d does not exist", localityID)
	}

	if err != nil {
		err = f.createLocality()
	}

	return
}

func (f *SellerFactory) createLocality() (err error) {
	localityFactory := NewLocalityFactory(f.db)
	_, err = localityFactory.Create(models.Locality{})

	return
}
