package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1SellerRoutes(db repository.RepoDB, rt *chi.Mux) {
	ct, err := initSellerController(db)
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/sellers", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.PostSeller)
		rt.Patch("/{id}", ct.PatchSeller)
		rt.Delete("/{id}", ct.Delete)
	})
}

func initSellerController(db repository.RepoDB) (ct controller.SellerDefault, err error) {
	sv := service.NewSellerService(db)

	ct = *controller.NewSellerController(sv)
	return
}
