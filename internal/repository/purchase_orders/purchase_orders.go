package purchase_orders_repository

import "database/sql"

type PurchaseOrdersRepository struct {
	DB *sql.DB
}

func NewPurchaseOrdersRepository(DB *sql.DB) *PurchaseOrdersRepository {
	repo := &PurchaseOrdersRepository{
		DB: DB,
	}
	return repo
}
