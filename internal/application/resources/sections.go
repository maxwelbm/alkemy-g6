package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitSections(db *sql.DB, router *chi.Mux) {
	rp := repository.NewSectionsRepository(db)
	// - service
	sv := service.NewSectionsService(rp)
	// - handler
	ct := controllers.NewSectionsController(sv)

	// - endpoints
	router.Route("/api/v1/sections", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetByID)
		rt.Get("/reportProducts", ct.ReportProducts)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}
