package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1BuyerRoutes(db repository.RepoDB, rt *chi.Mux) {
	ct, err := initBuyerController(db)
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/buyers", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.PostBuyer)
		rt.Patch("/{id}", ct.PatchBuyer)
		rt.Delete("/{id}", ct.DeleteBuyer)
	})
}

func initBuyerController(db repository.RepoDB) (ct controller.BuyerDefault, err error) {
	sv := service.NewBuyerService(db)

	ct = *controller.NewBuyerController(sv)
	return
}
