package repository

import (
	"database/sql"

	emp_models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	localities_repository "github.com/maxwelbm/alkemy-g6/internal/repository/localities"
	products_repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	sec_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	sellers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
	warehouse_repository "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
)

// Deprecated, create sql repositories for each model instead
type RepoDB struct {
	EmployeesDB emp_models.EmployeesRepository
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

func NewWarehousesRepository(DB *sql.DB) *warehouse_repository.WarehouseRepository {
	return &warehouse_repository.WarehouseRepository{
		DB: DB,
	}
}

func NewLocalityRepository(db *sql.DB) *localities_repository.LocalityRepository {
	return localities_repository.NewLocalityRepository(db)
}

func NewSellersRepository(db *sql.DB) *sellers_repository.SellersDefault {
	return sellers_repository.NewSellersRepository(db)
}

func NewSectionsRepository(DB *sql.DB) *sec_repository.SectionRepository {
	return &sec_repository.SectionRepository{
		DB: DB,
	}
}
