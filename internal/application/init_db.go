package application

import (
	"fmt"
	"os"

	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	emp_repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	prod_repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	sec_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	sel_repository "github.com/maxwelbm/alkemy-g6/internal/repository/seller"
)

func loadDB() (repo repository.RepoDB, err error) {
	emp, err := loadEmployeesRepository()
	if err != nil {
		return
	}
	prod, err := loadProductsRepository()
	if err != nil {
		return
	}
	sec, err := loadSectionsRepository()
	if err != nil {
		return
	}
	sell, err := loadSellersRepository()
	if err != nil {
		return
	}

	repo = repository.RepoDB{
		EmployeesDB: emp,
		ProductsDB:  prod,
		SectionsDB:  sec,
		SellersDB:   sell,
	}

	return
}

func loadEmployeesRepository() (repo *emp_repository.Employees, err error) {
	// loads employees from employees.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "employees.json")
	ld := loaders.NewEmployeesJSONFile(path)
	emp, err := ld.Load()
	if err != nil {
		return
	}

	repo = emp_repository.NewEmployees(emp)

	return
}

func loadProductsRepository() (repo *prod_repository.Products, err error) {
	// loads products from products.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "products.json")
	ld := loaders.NewProductJSONFile(path)
	prods, err := ld.Load()
	if err != nil {
		return
	}

	repo = prod_repository.NewProducts(prods)

	return
}

func loadSectionsRepository() (repo *sec_repository.Sections, err error) {
	// loads sections from sections.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "sections.json")
	ld := loaders.NewSectionJSONFile(path)
	sections, err := ld.Load()
	if err != nil {
		return
	}

	repo = sec_repository.NewSections(sections)

	return
}

func loadSellersRepository() (repo *sel_repository.SellerRepository, err error) {
	// loads sellers from sellers.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "sellers.json")
	ld := loaders.NewSellerJSONFile(path)
	sellers, err := ld.Load()
	if err != nil {
		return
	}

	repo = sel_repository.NewSellerRepository(sellers)

	return
}
