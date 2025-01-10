package resources

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func InitPurchaseOrders(db *sql.DB, rt *chi.Mux) {
	rp := repository.NewPurchaseOrdersRepository(db)
	// - service
	sv := service.NewPurchaseOrdersService(rp)
	// - handler
	ct := controllers.NewPurchaseOrdersController(sv)

	// - endpoints
	rt.Route("/api/v1/purchaseOrders", func(rt chi.Router) {
		rt.Post("/", ct.Create)
	})
}
