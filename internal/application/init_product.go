package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1ProductsRoutes(db repository.RepoDB, rt *chi.Mux) {
	ct, err := initProductsController(db)
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/products", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}

func initProductsController(db repository.RepoDB) (ct controller.ProductsDefault, err error) {
	sv := service.NewProductsDefault(db)

	ct = *controller.NewProductsDefault(sv)
	return
}
