package carries_repository

import (
	"database/sql"
)

type CarriesDefault struct {
	db *sql.DB
}

func NewCarriesRepository(db *sql.DB) *CarriesDefault {
	rp := &CarriesDefault{
		db: db,
	}
	return rp
}
