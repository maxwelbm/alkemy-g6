package repository

import (
	"database/sql"

	buy_models "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
	emp_models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	prod_models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	sell_models "github.com/maxwelbm/alkemy-g6/internal/models/seller"
	ware_models "github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
)

type RepoDB struct {
	BuyersDB    buy_models.BuyerRepository
	EmployeesDB emp_models.EmployeesRepository
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
