package carriesrp

import (
	"database/sql"
)

type CarriesDefault struct {
	db *sql.DB
}

func NewCarriesRepository(db *sql.DB) *CarriesDefault {
	return &CarriesDefault{
		db: db,
	}
}
