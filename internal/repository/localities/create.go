package localitiesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *LocalityRepository) Create(locDTO models.LocalityDTO) (loc models.Locality, err error) {
	// insert
	insertQuery := "INSERT INTO localities (locality_name, province_name, country_name) VALUES (?, ?, ?)"
	result, err := r.db.Exec(insertQuery, locDTO.LocalityName, locDTO.ProvinceName, locDTO.CountryName)

	if err != nil {
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return
	}

	getQuery := "SELECT id, locality_name, province_name, country_name FROM localities WHERE id = ?"
	err = r.db.QueryRow(getQuery, lastInsertID).Scan(&loc.ID, &loc.LocalityName, &loc.ProvinceName, &loc.CountryName)

	return
}
