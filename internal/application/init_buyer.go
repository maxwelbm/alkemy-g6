package application

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	buyerController "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	buyerRepository "github.com/maxwelbm/alkemy-g6/internal/repository/buyer"
	buyerService "github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1BuyerRoutes(rt *chi.Mux) {
	ct, err := initBuyerController()
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

func initBuyerController() (ct buyerController.BuyerDefault, err error) {
	repo, err := loadBuyersRepository()
	if err != nil {
		return
	}
	sv := buyerService.NewBuyerService(repo)

	ct = *buyerController.NewBuyerController(sv)
	return
}

func loadBuyersRepository() (repo *buyerRepository.BuyerRepository, err error) {
	// loads products from products.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "buyers.json")
	ld := loaders.NewBuyerJSONFile(path)
	buyers, err := ld.Load()
	if err != nil {
		return
	}

	repo = buyerRepository.NewBuyerRepository(buyers)

	return
}
