package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *Employees) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	employees, err := c.sv.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	list := EmployeesAttributes{
		ID:           employees.ID,
		CardNumberID: employees.CardNumberID,
		FirstName:    employees.FirstName,
		LastName:     employees.LastName,
		WarehouseID:  employees.WarehouseID,
	}

	responseEmp := EmployeesFinalJSON{
		Data: list,
	}

	response.JSON(w, http.StatusOK, responseEmp)
}
