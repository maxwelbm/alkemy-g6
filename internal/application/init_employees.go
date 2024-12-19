package application

import (
	"log"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1EmployeesRoutes(db repository.RepoDB, rt *chi.Mux) {
	ct, err := initEmployeesController(db)
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/employees", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetByID)
	})
}

func initEmployeesController(db repository.RepoDB) (ct controller.Employees, err error) {
	sv := *service.NewEmployeesDefault(db)

	ct = *controller.NewEmployeesDefault(sv)
	return
}
