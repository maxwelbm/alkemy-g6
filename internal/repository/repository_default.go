package repository

import (
	"database/sql"

	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	carries_repository "github.com/maxwelbm/alkemy-g6/internal/repository/carries"
	employees_repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	inbound_orders_repository "github.com/maxwelbm/alkemy-g6/internal/repository/inbound_orders"
	localities_repository "github.com/maxwelbm/alkemy-g6/internal/repository/localities"
	product_batches_repository "github.com/maxwelbm/alkemy-g6/internal/repository/product_batches"
	product_records_repository "github.com/maxwelbm/alkemy-g6/internal/repository/product_records"
	products_repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	purchase_orders_repository "github.com/maxwelbm/alkemy-g6/internal/repository/purchase_orders"
	sections_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	sellers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
	warehouse_repository "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
)

func NewBuyersRepository(db *sql.DB) *buyers_repository.BuyerRepository {
	return buyers_repository.NewBuyersRepository(db)
}

func NewProductsRepository(DB *sql.DB) *products_repository.Products {
	return &products_repository.Products{
		DB: DB,
	}
}

func NewWarehousesRepository(db *sql.DB) *warehouse_repository.WarehouseRepository {
	return warehouse_repository.NewWarehouseRepository(db)
}

func NewLocalityRepository(db *sql.DB) *localities_repository.LocalityRepository {
	return localities_repository.NewLocalityRepository(db)
}

func NewSellersRepository(db *sql.DB) *sellers_repository.SellersDefault {
	return sellers_repository.NewSellersRepository(db)
}

func NewCarriesRepository(db *sql.DB) *carries_repository.CarriesDefault {
	return carries_repository.NewCarriesRepository(db)
}

func NewSectionsRepository(db *sql.DB) *sections_repository.SectionRepository {
	return sections_repository.NewSectionsRepository(db)
}

func NewProductBatchesRepository(db *sql.DB) *product_batches_repository.ProductBatchesRepository {
	return product_batches_repository.NewProductBatchesRepository(db)
}

func NewProductRecordsRepository(db *sql.DB) *product_records_repository.ProductRecordsDefault {
	return product_records_repository.NewProductRecordsRepository(db)
}

func NewEmployeesRepository(db *sql.DB) *employees_repository.EmployeesRepository {
	return employees_repository.NewEmployeesRepository(db)
}

func NewPurchaseOrdersRepository(db *sql.DB) *purchase_orders_repository.PurchaseOrdersRepository {
	return purchase_orders_repository.NewPurchaseOrdersRepository(db)
}

func NewInboundOrdersRepository(db *sql.DB) *inbound_orders_repository.InboundOrdersRepository {
	return inbound_orders_repository.NewInboundOrdersRepository(db)
}
