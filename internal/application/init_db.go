package application

import (
	"fmt"
	"os"

	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	emp_repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	sec_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	war_repository "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
)

func loadDB() (repo repository.RepoDB, err error) {
	emp, err := loadEmployeesRepository()
	if err != nil {
		return
	}
	sec, err := loadSectionsRepository()
	if err != nil {
		return
	}
	ware, err := loadWarehousesRepository()
	if err != nil {
		return
	}

	repo = repository.RepoDB{
		EmployeesDB: emp,
		SectionsDB:  sec,
		WarehouseDB: ware,
	}

	return
}

func loadEmployeesRepository() (repo *emp_repository.Employees, err error) {
	// loads employees from employees.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "employees.json")
	ld := loaders.NewEmployeesJSONFile(path)
	emp, err := ld.Load()
	if err != nil {
		return
	}

	repo = emp_repository.NewEmployees(emp)

	return
}

func loadWarehousesRepository() (repo *war_repository.Warehouses, err error) {
	// loads warehouses from warehouses.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "warehouses.json")
	ld := loaders.NewWarehouseJSONFile(path)
	warehouses, err := ld.Load()
	if err != nil {
		return
	}

	repo = war_repository.NewWarehouses(warehouses)

	return
}

func loadSectionsRepository() (repo *sec_repository.Sections, err error) {
	// loads sections from sections.json file
	path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "sections.json")
	ld := loaders.NewSectionJSONFile(path)
	sections, err := ld.Load()
	if err != nil {
		return
	}

	repo = sec_repository.NewSections(sections)

	return
}
