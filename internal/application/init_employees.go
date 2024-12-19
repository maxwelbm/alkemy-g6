package application

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	controller "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/internal/service"
)

func buildApiV1EmployeesRoutes(rt *chi.Mux) {
	ct, err := initEmployeesController()
	if err != nil {
		log.Fatal(err)
	}

	rt.Route("/api/v1/employees", func(rt chi.Router) {
		rt.Get("/", ct.GetAll)
		rt.Get("/{id}", ct.GetByID)
		rt.Post("/", ct.Create)
		rt.Patch("/{id}", ct.Update)
		rt.Delete("/{id}", ct.Delete)
	})
}

func initEmployeesController() (ct controller.Employees, err error) {
	rp, err := loadEmployeesRepository()
	if err != nil {
		return
	}
	sv := *service.NewEmployeesDefault(&rp)

	ct = *controller.NewEmployeesDefault(&sv)
	return
}

func loadEmployeesRepository() (repo repository.Employees, err error) {
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "employees.json")
	ld := loaders.NewEmployeesJSONFile(path)
	emp, err := ld.Load()
	if err != nil {
		return
	}

	repo = *repository.NewEmployees(emp)

	return
}
