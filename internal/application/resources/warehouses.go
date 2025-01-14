package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitWarehouses(db *sql.DB, rt *chi.Mux) {
	rp := repository.NewWarehousesRepository(db)
	// - service
	sv := service.NewWarehousesService(rp)
	// - handler
	ct := controllers.NewWarehousesController(sv)

	// - endpoints
	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetByID)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}
