package application

import (
	"fmt"
	"log"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	"github.com/maxwelbm/alkemy-g6/internal/repository/warehouse"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1WarehousesRoutes(rt *chi.Mux) {
	ct, err := initWarehousesController()
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

func initWarehousesController() (ct controller.WarehouseDefault, err error) {
	repo, err := loadWarehousesRepository()
	if err != nil {
		return
	}
	sv := service.NewWarehouseDefault(repo)

	ct = *controller.NewWarehouseDefault(sv)
	return
}

func loadWarehousesRepository() (repo *repository.Warehouses, err error) {
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "warehouses.json")
	ld := loaders.NewWarehouseJSONFile(path)
	prods, err := ld.Load()
	if err != nil {
		return
	}

	repo = repository.NewWarehouses(prods)

	return
}