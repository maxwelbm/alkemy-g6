package application

import (
	"fmt"
	"os"

	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	emp_repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
)

func loadDB() (repo repository.RepoDB, err error) {
	emp, err := loadEmployeesRepository()
	if err != nil {
		return
	}

	repo = repository.RepoDB{
		EmployeesDB: emp,
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
