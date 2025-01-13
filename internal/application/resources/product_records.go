package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitProductRecords(db *sql.DB, router *chi.Mux) {
	rp := repository.NewProductRecordsRepository(db)
	// - service
	sv := service.NewProductRecordsService(rp)
	// - handler
	ct := controllers.NewProductRecordsController(sv)

	// - endpoints
	router.Route("/api/v1/productRecords", func(rt chi.Router) {
		rt.Post("/", ct.Create)
	})
}
