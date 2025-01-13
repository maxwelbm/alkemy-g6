package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitInboundOrders(db *sql.DB, router *chi.Mux) {
	rp := repository.NewInboundOrdersRepository(db)
	// - service
	sv := service.NewInboundOrdersService(rp)
	// - handler
	ct := controllers.NewInboundOrdersController(sv)

	// - endpoints
	router.Route("/api/v1/inboundOrders", func(rt chi.Router) {
		rt.Post("/", ct.Create)
	})
}
