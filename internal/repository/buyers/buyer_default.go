package buyers_repository

import (
	"database/sql"
)

type BuyerRepository struct {
	DB *sql.DB
}

func NewBuyersRepository(DB *sql.DB) *BuyerRepository {
	repo := &BuyerRepository{
		DB: DB,
	}
	return repo
}
