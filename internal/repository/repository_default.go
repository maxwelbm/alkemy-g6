package repository

import (
	"database/sql"

	emp_models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	prod_models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	ware_models "github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	sellers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
)

type RepoDB struct {
	EmployeesDB emp_models.EmployeesRepository
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
