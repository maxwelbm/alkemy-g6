package repository

import (
	emp_models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	prod_models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	sec_models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	sell_models "github.com/maxwelbm/alkemy-g6/internal/models/seller"
	ware_models "github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

type RepoDB struct {
	EmployeesDB emp_models.EmployeesRepository
	ProductsDB  prod_models.ProductRepository
	SectionsDB  sec_models.SectionRepository
	SellersDB   sell_models.SellerRepository
	WarehouseDB ware_models.WarehouseRepository
}
