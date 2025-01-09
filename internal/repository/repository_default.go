package repository

import (
	"database/sql"

	prod_models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	sell_models "github.com/maxwelbm/alkemy-g6/internal/models/seller"
	ware_models "github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	employees_repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
)

type RepoDB struct {
	ProductsDB  prod_models.ProductRepository
	SectionsDB  sec_models.SectionRepository
	SellersDB   sell_models.SellerRepository
	WarehouseDB ware_models.WarehouseRepository
}

func NewBuyersRepository(DB *sql.DB) *buyers_repository.BuyerRepository {
	return &buyers_repository.BuyerRepository{
		DB: DB,
	}
}

func NewEmployeesRepository(DB *sql.DB) *employees_repository.EmployeesRepository {
	return &employees_repository.EmployeesRepository{
		DB: DB,
	}
}
