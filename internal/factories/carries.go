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

func (f *CarryFactory) Build(carry models.Carry) models.Carry {
	populateCarryParams(&carry)

	return carry
}

func (f *CarryFactory) Create(carry models.Carry) (record models.Carry, err error) {
	populateCarryParams(&carry)

	if err = f.checkLocalityExists(carry.LocalityID); err != nil {
		return carry, err
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

	switch carry.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(carry.ID)+",")
	}

	_, err = f.db.Exec(query,
		carry.CID,
		carry.CompanyName,
		carry.Address,
		carry.PhoneNumber,
		carry.LocalityID,
	)

	return carry, err
}

func populateCarryParams(carry *models.Carry) {
	defaultCarry := defaultCarry()
	if carry == nil {
		carry = &defaultCarry
	}

	if carry.CID == "" {
		carry.CID = defaultCarry.CID
	}

	if carry.CompanyName == "" {
		carry.CompanyName = defaultCarry.CompanyName
	}

	if carry.Address == "" {
		carry.Address = defaultCarry.Address
	}

	if carry.PhoneNumber == "" {
		carry.PhoneNumber = defaultCarry.PhoneNumber
	}

	if carry.LocalityID == 0 {
		carry.LocalityID = defaultCarry.LocalityID
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
