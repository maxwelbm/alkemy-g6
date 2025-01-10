package sellers_repository

import (
	"database/sql"
)

type SellersDefault struct {
	db *sql.DB
}

func NewSellersRepository(db *sql.DB) *SellersDefault {
	rp := &SellersDefault{
		db: db,
	}
	return rp
}
