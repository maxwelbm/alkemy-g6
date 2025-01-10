package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitCarries(db *sql.DB, router *chi.Mux) {
	rp := repository.NewCarriesRepository(db)
	// - service
	sv := service.NewCarriesService(rp)
	// - handler
	ct := controllers.NewCarriesController(sv)

	// - endpoints
	router.Route("/api/v1/carries", func(rt chi.Router) {
		rt.Post("/", ct.Create)
	})
}
