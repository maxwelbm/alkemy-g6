package buyersrp

import (
	"database/sql"
)

type BuyerRepository struct {
	db *sql.DB
}

func NewBuyersRepository(db *sql.DB) *BuyerRepository {
	return &BuyerRepository{
		db: db,
	}
}
