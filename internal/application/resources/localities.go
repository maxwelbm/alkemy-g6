package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitLocalities(db *sql.DB, router *chi.Mux) {
	rp := repository.NewLocalityRepository(db)
	// - service
	sv := service.NewLocalitiesService(rp)
	// - handler
	ct := controllers.NewLocalityController(sv)

	// - endpoints
	router.Route("/api/v1/localities", func(r chi.Router) {
		r.Get("/reportSellers", ct.ReportSellers)
		r.Get("/reportCarries", ct.ReportCarries)
		r.Post("/", ct.Create)
	})
}
