package buyersrp

import (
	"database/sql"
)

type BuyerRepository struct {
	db *sql.DB
}

func NewBuyersRepository(db *sql.DB) *BuyerRepository {
	repo := &BuyerRepository{
		db: db,
	}
	return repo
}
