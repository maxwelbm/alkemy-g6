package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"

	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitSellers(db *sql.DB, router *chi.Mux) {
	rp := repository.NewSellersRepository(db)
	// - service
	sv := service.NewSellersService(rp)
	// - handler
	ct := controllers.NewSellersController(sv)

	// - endpoints
	router.Route("/api/v1/sellers", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetById)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}
