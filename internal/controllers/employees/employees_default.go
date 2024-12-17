package employees_controller

import (
	"github.com/maxwelbm/alkemy-g6b/internal/models"
)

type Employees struct {
	sv models.EmployeesService
}

func NewEmployeesDefault(sv models.EmployeesService) *Employees {
	return &Employees{sv: sv}
}
