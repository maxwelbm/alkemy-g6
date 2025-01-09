package repository

import (
	"database/sql"

	emp_models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	prod_models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	warehouse_repository "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
	sellers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
)

// Deprecated, create sql repositories for each model instead
type RepoDB struct {
	EmployeesDB emp_models.EmployeesRepository
	ProductsDB  prod_models.ProductRepository
	SectionsDB  sec_models.SectionRepository
}

func NewBuyersRepository(DB *sql.DB) *buyers_repository.BuyerRepository {
	return &buyers_repository.BuyerRepository{
		DB: DB,
	}
}

func NewWarehousesRepository(DB *sql.DB) *warehouse_repository.WarehouseRepository {
	return &warehouse_repository.WarehouseRepository{
		DB: DB,
	}
}

func NewSellersRepository(DB *sql.DB) *sellers_repository.SellersDefault {
	return &sellers_repository.SellersDefault{
		DB: DB,
	}
}
