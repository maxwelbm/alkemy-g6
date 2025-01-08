package repository

import (
	"database/sql"

	prod_models "github.com/maxwelbm/alkemy-g6/internal/models" //remove this when sellers is done
	emp_models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	sell_models "github.com/maxwelbm/alkemy-g6/internal/models/seller"
	ware_models "github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	products_repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
)

type RepoDB struct {
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

func NewProductsRepository(DB *sql.DB) *products_repository.Products {
	return &products_repository.Products{
		DB: DB,
	}
}
