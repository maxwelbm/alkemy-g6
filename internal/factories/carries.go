package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type CarryFactory struct {
	db *sql.DB
}

func NewCarryFactory(db *sql.DB) *CarryFactory {
	return &CarryFactory{db: db}
}

func defaultCarry() models.Carry {
	return models.Carry{
		CID:         randstr.Alphanumeric(8),
		CompanyName: randstr.Chars(8),
		Address:     randstr.Chars(8),
		PhoneNumber: randstr.Chars(11),
		LocalityID:  1,
	}
}

func (f *CarryFactory) Create(carrie models.Carry) (record models.Carry, err error) {
	populateCarryParams(&carrie)

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

func populateCarryParams(carrie *models.Carry) {
	defaultCarry := defaultCarry()
	if carrie == nil {
		carrie = &defaultCarry
	}

	if carrie.CID == "" {
		carrie.CID = defaultCarry.CID
	}

	if carrie.CompanyName == "" {
		carrie.CompanyName = defaultCarry.CompanyName
	}

	if carrie.Address == "" {
		carrie.Address = defaultCarry.Address
	}

	if carrie.PhoneNumber == "" {
		carrie.PhoneNumber = defaultCarry.PhoneNumber
	}

	if carrie.LocalityID == 0 {
		carrie.LocalityID = defaultCarry.LocalityID
	}
}

func (f *CarryFactory) checkLocalityExists(localityID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM localities WHERE id = ?`, localityID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createLocality()

	return
}

func (f *CarryFactory) createLocality() (err error) {
	localityFactory := NewLocalityFactory(f.db)
	_, err = localityFactory.Create(models.Locality{})

	return
}
