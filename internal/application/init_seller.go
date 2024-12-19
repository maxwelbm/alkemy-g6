package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	sellerController "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	service "github.com/maxwelbm/alkemy-g6/internal/service"
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

func initSellerController(db repository.RepoDB) (ct sellerController.SellerDefault, err error) {
	sv := service.NewSellerService(db)

	ct = *sellerController.NewSellerController(sv)
	return
}
