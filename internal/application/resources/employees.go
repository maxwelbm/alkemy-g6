package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitEmployees(db *sql.DB, router *chi.Mux) {
	rp := repository.NewEmployeesRepository(db)
	// - service
	sv := service.NewEmployeesService(rp)
	// - handler
	ct := controllers.NewEmployeesController(sv)

	// - endpoints
	router.Route("/api/v1/employees", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetByID)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}
