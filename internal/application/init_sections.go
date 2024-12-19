package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1SectionsRoutes(db repository.RepoDB, rt *chi.Mux) {
	ct, err := initSectionsController(db)
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/sections", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.Create)
	})
}

func initSectionsController(db repository.RepoDB) (ct controller.SectionsDefault, err error) {
	sv := service.NewSectionDefault(db)

	ct = *controller.NewSectionsDefault(sv)
	return
}
