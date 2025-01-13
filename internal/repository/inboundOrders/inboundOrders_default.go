package repository

import (
	"database/sql"
	"errors"
)

var (
	ErrInboundOrdersRepositoryNotFound = errors.New("Inbound Orders not found")
)

type InboundOrdersRepository struct {
	DB *sql.DB
}

func NewInboundOrdersRepository(DB *sql.DB) *InboundOrdersRepository {
	repo := &InboundOrdersRepository{
		DB: DB,
	}
	return repo
}
