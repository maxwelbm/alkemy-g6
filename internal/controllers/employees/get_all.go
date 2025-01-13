package employees_controller

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *EmployeesController) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := c.SV.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var list []EmployeesAttributes
	for _, value := range employees {
		list = append(list, EmployeesAttributes{
			ID:           value.ID,
			CardNumberID: value.CardNumberID,
			FirstName:    value.FirstName,
			LastName:     value.LastName,
			WarehouseID:  value.WarehouseID,
		})
	}

	responseEmp := EmployeesFinalJSON{
		Data: list,
	}

	response.JSON(w, http.StatusOK, responseEmp)
}
