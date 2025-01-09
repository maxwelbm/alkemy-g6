package repository

import (
	"database/sql"

	emp_models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	prod_models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	buyers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	localities_repository "github.com/maxwelbm/alkemy-g6/internal/repository/localities"
	sellers_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
	warehouse_repository "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
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

func NewLocalityRepository(db *sql.DB) *localities_repository.LocalityRepository {
	return localities_repository.NewLocalityRepository(db)
}

func NewSellersRepository(DB *sql.DB) *sellers_repository.SellersDefault {
	return &sellers_repository.SellersDefault{
		DB: DB,
	}
}
