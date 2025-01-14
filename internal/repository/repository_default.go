package repository

import (
	"database/sql"

	buyersrp "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	carriesrp "github.com/maxwelbm/alkemy-g6/internal/repository/carries"
	employeesrp "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	inboundordersrp "github.com/maxwelbm/alkemy-g6/internal/repository/inbound_orders"
	localitiesrp "github.com/maxwelbm/alkemy-g6/internal/repository/localities"
	productbatchesrp "github.com/maxwelbm/alkemy-g6/internal/repository/product_batches"
	productrecordsrp "github.com/maxwelbm/alkemy-g6/internal/repository/product_records"
	productsrp "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	purchaseordersrp "github.com/maxwelbm/alkemy-g6/internal/repository/purchase_orders"
	sectionsrp "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	sellersrp "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
	warehouserp "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
)

func NewBuyersRepository(db *sql.DB) *buyersrp.BuyerRepository {
	return buyersrp.NewBuyersRepository(db)
}

func NewProductsRepository(DB *sql.DB) *productsrp.Products {
	return &productsrp.Products{
		DB: DB,
	}
}

func NewWarehousesRepository(db *sql.DB) *warehouserp.WarehouseRepository {
	return warehouserp.NewWarehouseRepository(db)
}

func NewLocalityRepository(db *sql.DB) *localitiesrp.LocalityRepository {
	return localitiesrp.NewLocalityRepository(db)
}

func NewSellersRepository(db *sql.DB) *sellersrp.SellersDefault {
	return sellersrp.NewSellersRepository(db)
}

func NewCarriesRepository(db *sql.DB) *carriesrp.CarriesDefault {
	return carriesrp.NewCarriesRepository(db)
}

func NewSectionsRepository(db *sql.DB) *sectionsrp.SectionRepository {
	return sectionsrp.NewSectionsRepository(db)
}

func NewProductBatchesRepository(db *sql.DB) *productbatchesrp.ProductBatchesRepository {
	return productbatchesrp.NewProductBatchesRepository(db)
}

func NewProductRecordsRepository(db *sql.DB) *productrecordsrp.ProductRecordsDefault {
	return productrecordsrp.NewProductRecordsRepository(db)
}

func NewEmployeesRepository(db *sql.DB) *employeesrp.EmployeesRepository {
	return employeesrp.NewEmployeesRepository(db)
}

func NewPurchaseOrdersRepository(db *sql.DB) *purchaseordersrp.PurchaseOrdersRepository {
	return purchaseordersrp.NewPurchaseOrdersRepository(db)
}

func NewInboundOrdersRepository(db *sql.DB) *inboundordersrp.InboundOrdersRepository {
	return inboundordersrp.NewInboundOrdersRepository(db)
}
