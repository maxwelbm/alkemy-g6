package application

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	sellerController "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	sellerRepository "github.com/maxwelbm/alkemy-g6/internal/repository/seller"
	sellerService "github.com/maxwelbm/alkemy-g6/internal/service/seller"
)

func buildApiV1SellerRoutes(rt *chi.Mux) {
	ct, err := initSellerController()
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/sellers", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		//		rt.Patch("/{id}", ct.PatchSeller)
		rt.Delete("/{id}", ct.Delete)
	})
}

func initSellerController() (ct sellerController.SellerDefault, err error) {
	repo, err := loadSellerRepository()
	if err != nil {
		return
	}
	sv := sellerService.NewSellerService(repo)

	ct = *sellerController.NewSellerController(sv)
	return
}

func loadSellerRepository() (repo *sellerRepository.SellerRepository, err error) {
	// loads products from products.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "sellers.json")
	ld := loaders.NewSellerJSONFile(path)
	sellers, err := ld.Load()
	if err != nil {
		return
	}

	repo = sellerRepository.NewSellerRepository(sellers)

	return
}
