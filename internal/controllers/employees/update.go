package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *Employees) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}
	var employeesJSON EmployeesReqJSON
	err = json.NewDecoder(r.Body).Decode(&employeesJSON)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Failed to update employees")
		return
	}

	var employees models.EmployeesDTO
	employees.ID = id
	if employeesJSON.CardNumberID != nil {
		employees.CardNumberID = *employeesJSON.CardNumberID
	}
	if employeesJSON.FirstName != nil {
		employees.FirstName = *employeesJSON.FirstName
	}
	if employeesJSON.LastName != nil {
		employees.LastName = *employeesJSON.LastName
	}
	if employeesJSON.WarehouseID != nil {

		employees.WarehouseID = *employeesJSON.WarehouseID
		if employees.WarehouseID <= 0 {
			response.Error(w, http.StatusBadRequest, "WarehouseID invalid")
			return
		}
	}

	emp, err := c.sv.Update(employees, id)
	if errors.Is(err, repository.ErrEmployeesRepositoryDuplicatedCode) {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}
	if errors.Is(err, repository.ErrEmployeesRepositoryNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if errors.Is(err, service.ErrWareHousesServiceNotFound) {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data := EmployeesResJSON{
		Message: "Sucess created",
		Data: EmployeesAttributes{
			ID:           emp.ID,
			CardNumberID: emp.CardNumberID,
			FirstName:    emp.FirstName,
			LastName:     emp.LastName,
			WarehouseID:  emp.WarehouseID,
		},
	}

	response.JSON(w, http.StatusOK, data)
}
