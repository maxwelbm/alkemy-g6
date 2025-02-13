package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type CarrieFactory struct {
	db *sql.DB
}

func NewCarrieFactory(db *sql.DB) *CarrieFactory {
	return &CarrieFactory{db: db}
}

func defaultCarrie() models.Carry {
	return models.Carry{
		CID:         RandAlphanumeric(8),
		CompanyName: RandChars(8),
		Address:     RandChars(8),
		PhoneNumber: RandChars(11),
		LocalityID:  1,
	}
}

func (f *CarrieFactory) Create(carrie models.Carry) (record models.Carry, err error) {
	populateCarrieParams(&carrie)

	if err = f.checkLocalityExists(carrie.LocalityID); err != nil {
		return carrie, err
	}

	query := `
		INSERT INTO carries 
			(
			%s
			cid,
			company_name, 
			address, 
			phone_number, 
			locality_id
			) 
		VALUES (%s?, ?, ?, ?, ?)
	`

	switch carrie.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(carrie.ID)+",")
	}

	_, err = f.db.Exec(query,
		carrie.CID,
		carrie.CompanyName,
		carrie.Address,
		carrie.PhoneNumber,
		carrie.LocalityID,
	)

	return carrie, err
}

func populateCarrieParams(carrie *models.Carry) {
	defaultCarrie := defaultCarrie()
	if carrie == nil {
		carrie = &defaultCarrie
	}

	if carrie.CID == "" {
		carrie.CID = defaultCarrie.CID
	}

	if carrie.CompanyName == "" {
		carrie.CompanyName = defaultCarrie.CompanyName
	}

	if carrie.Address == "" {
		carrie.Address = defaultCarrie.Address
	}

	if carrie.PhoneNumber == "" {
		carrie.PhoneNumber = defaultCarrie.PhoneNumber
	}

	if carrie.LocalityID == 0 {
		carrie.LocalityID = defaultCarrie.LocalityID
	}
}

func (f *CarrieFactory) checkLocalityExists(localityID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM localities WHERE id = ?`, localityID).Scan(&count)

	if err != nil {
		return
	}

	if count == 0 {
		err = fmt.Errorf("carrie with id %d does not exist", localityID)
	}

	if err != nil {
		err = f.createLocality()
	}

	return
}

func (f *CarrieFactory) createLocality() (err error) {
	localityFactory := NewLocalityFactory(f.db)
	_, err = localityFactory.Create(models.Locality{})

	return
}
