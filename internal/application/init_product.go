package application

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	products_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	product_repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1ProductsRoutes(rt *chi.Mux) {
	ct, err := initProductsController()
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/products", func(rt chi.Router) {
		rt.Get("/", ct.GetAll())
	})
}

func initProductsController() (ct products_controller.ProductsDefault, err error) {
	repo, err := loadProductsRepository()
	if err != nil {
		return
	}
	sv := service.NewProductsDefault(repo)

	ct = *products_controller.NewProductsDefault(sv)
	return
}

func loadProductsRepository() (repo product_repository.Products, err error) {
	// loads products from products.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "products.json")
	ld := loaders.NewProductJSONFile(path)
	prods, err := ld.Load()
	if err != nil {
		return
	}

	repo = *product_repository.NewProducts(prods)

	return
}
