package application

import (
	"fmt"
	"os"

	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	prod_repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	sec_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	sel_repository "github.com/maxwelbm/alkemy-g6/internal/repository/seller"
	war_repository "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
)

func loadDB() (repo repository.RepoDB, err error) {
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
	ware, err := loadWarehousesRepository()
	if err != nil {
		return
	}

	repo = repository.RepoDB{
		ProductsDB:  prod,
		SectionsDB:  sec,
		SellersDB:   sell,
		WarehouseDB: ware,
	}

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

func loadWarehousesRepository() (repo *war_repository.Warehouses, err error) {
	// loads warehouses from warehouses.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "warehouses.json")
	ld := loaders.NewWarehouseJSONFile(path)
	warehouses, err := ld.Load()
	if err != nil {
		return
	}

	repo = war_repository.NewWarehouses(warehouses)

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
