package repository

import (
	"database/sql"
)

type Products struct {
	DB *sql.DB
}

func NewProducts(db *sql.DB) *Products {
	repo := &Products{
		DB: db,
	}
	
	return repo
}