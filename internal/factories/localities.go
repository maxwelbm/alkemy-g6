package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type LocalityFactory struct {
	db *sql.DB
}

func NewLocalityFactory(db *sql.DB) *LocalityFactory {
	return &LocalityFactory{db: db}
}

func defaultLocality() models.Locality {
	return models.Locality{
		LocalityName: RandChars(8),
		ProvinceName: RandChars(8),
		CountryName:  RandChars(8),
	}
}

func (f LocalityFactory) Build(locality models.Locality) models.Locality {
	populateLocalityParams(&locality)

	return locality
}

func (f *LocalityFactory) Create(locality models.Locality) (record models.Locality, err error) {
	populateLocalityParams(&locality)

	query := `
		INSERT INTO localities 
			(
			%s
			locality_name, 
			province_name, 
			country_name
			) 
		VALUES (%s?, ?, ?)
	`

	switch locality.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(locality.ID)+",")
	}

	result, err := f.db.Exec(query,
		locality.LocalityName,
		locality.ProvinceName,
		locality.CountryName,
	)

	if err != nil {
		return record, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return record, err
	}

	locality.ID = int(id)

	return locality, err
}

func populateLocalityParams(locality *models.Locality) {
	defaultLocality := defaultLocality()
	if locality == nil {
		locality = &defaultLocality
	}

	if locality.LocalityName == "" {
		locality.LocalityName = defaultLocality.LocalityName
	}

	if locality.ProvinceName == "" {
		locality.ProvinceName = defaultLocality.ProvinceName
	}

	if locality.CountryName == "" {
		locality.CountryName = defaultLocality.CountryName
	}
}
