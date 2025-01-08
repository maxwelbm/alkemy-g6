package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitProducts(db *sql.DB, router *chi.Mux) {
	rp := repository.NewProductsRepository(db)
	// - service
	sv := service.NewProductsService(rp)
	// - handler
	ct := controllers.NewProductsController(sv)

	// - endpoints
	router.Route("/api/v1/products", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}
