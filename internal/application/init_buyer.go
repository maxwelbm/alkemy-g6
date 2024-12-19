package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	buyerController "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	buyerService "github.com/maxwelbm/alkemy-g6/internal/service"
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

func initBuyerController(db repository.RepoDB) (ct buyerController.BuyerDefault, err error) {
	sv := buyerService.NewBuyerService(db)

	ct = *buyerController.NewBuyerController(sv)
	return
}
