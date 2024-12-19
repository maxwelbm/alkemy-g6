package application

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1SectionsRoutes(rt *chi.Mux) {
	ct, err := initSectionsController()
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/sections", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}

func initSectionsController() (ct controller.SectionsDefault, err error) {
	repo, err := loadSectionsRepository()
	if err != nil {
		return
	}
	sv := service.NewSectionDefault(&repo)

	ct = *controller.NewSectionsDefault(sv)
	return
}

func loadSectionsRepository() (repo repository.Sections, err error) {
	// loads products from products.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "sections.json")
	ld := loaders.NewSectionJSONFile(path)
	sections, err := ld.Load()
	if err != nil {
		return
	}

	repo = *repository.NewSections(sections)

	return
}
