package employees_controller

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g6b/internal/loaders"
)

func (c *Employees) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := c.sv.GetAll()
	if err != nil {
		return
	}

	data := make(map[int]loaders.EmployeesJSON)
	for key, value := range employees {
		data[key] = loaders.EmployeesJSON{
			ID:           value.ID,
			CardNumberID: value.CardNumberID,
			FirstName:    value.FirstName,
			LastName:     value.LastName,
			WarehouseID:  value.WarehouseID,
		}
	}

	response.JSON

	return
}
