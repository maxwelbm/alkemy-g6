package inboundordersrp

import (
	"database/sql"
)

type InboundOrdersRepository struct {
	db *sql.DB
}

func NewInboundOrdersRepository(db *sql.DB) *InboundOrdersRepository {
	return &InboundOrdersRepository{
		db: db,
	}
}
