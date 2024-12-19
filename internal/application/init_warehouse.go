package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1WarehousesRoutes(db repository.RepoDB, rt *chi.Mux) {
	ct, err := initWarehousesController(db)
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
	})
}

func initWarehousesController(db repository.RepoDB) (ct controller.WarehouseDefault, err error) {
	sv := service.NewWarehouseDefault(db)

	ct = *controller.NewWarehouseDefault(sv)
	return
}
