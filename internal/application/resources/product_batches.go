package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitProductBatches(db *sql.DB, router *chi.Mux) {
	rp := repository.NewProductBatchesRepository(db)
	// - service
	sv := service.NewProductBatchesDefault(rp)
	// - handler
	ct := controllers.NewProductBatchesController(sv)

	// - endpoints
	router.Route("/api/v1/productBatches", func(rt chi.Router) {
		rt.Post("/", ct.Create)
	})
}
