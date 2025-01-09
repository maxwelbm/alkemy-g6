package sellers_repository

import (
	"database/sql"
)

type SellersDefault struct {
	DB *sql.DB
}

func NewSellersRepository(DB *sql.DB) *SellersDefault {
	rp := &SellersDefault{
		DB: DB,
	}
	return rp
}
