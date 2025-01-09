package repository

import (
	"database/sql"

	prod_models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	ware_models "github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	employees_repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	sellers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
)

// Deprecated, create sql repositories for each model instead
type RepoDB struct {
	ProductsDB  prod_models.ProductRepository
	SectionsDB  sec_models.SectionRepository
	WarehouseDB ware_models.WarehouseRepository
}

func NewBuyersRepository(DB *sql.DB) *buyers_repository.BuyerRepository {
	return &buyers_repository.BuyerRepository{
		DB: DB,
	}
}

func NewSellersRepository(DB *sql.DB) *sellers_repository.SellersDefault {
	return &sellers_repository.SellersDefault{
		DB: DB,
	}
}

func NewEmployeesRepository(DB *sql.DB) *employees_repository.EmployeesRepository {
	return &employees_repository.EmployeesRepository{
		DB: DB,
	}
}
