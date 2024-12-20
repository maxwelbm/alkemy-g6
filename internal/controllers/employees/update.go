package controller

import (
	"encoding/json"
	"errors"
	"fmt"
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

	err = validateNewEmployees(employeesJSON)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	newEmployees := models.EmployeesDTO{
		ID:           employeesJSON.ID,
		CardNumberID: employeesJSON.CardNumberID,
		FirstName:    employeesJSON.FirstName,
		LastName:     employeesJSON.LastName,
		WarehouseID:  employeesJSON.WarehouseID,
	}

	emp, err := c.sv.Update(newEmployees, id)
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
		Message: "Sucess",
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

func validateUpdateEmployees(e EmployeesReqJSON) (err error) {

	var errosEmp []string
	if e.CardNumberID == nil {
		errosEmp = append(errosEmp, "CardNumberID cannot be nil")
	} else if *e.CardNumberID == "" {
		errosEmp = append(errosEmp, "CardNumberID cannot be empty")
	}

	// FirstName
	if e.FirstName == nil {
		errosEmp = append(errosEmp, "FirstName cannot be nil")
	} else if *e.CardNumberID == "" {
		errosEmp = append(errosEmp, "FirstName cannot be empty")
	}

	// LastName
	if e.LastName == nil {
		errosEmp = append(errosEmp, "LastName cannot be empty")
	} else if *e.CardNumberID == "" {
		errosEmp = append(errosEmp, "LastName cannot be empty")
	}

	// WarehouseID
	if e.WarehouseID != nil {
		if *e.WarehouseID <= 0 {
			errosEmp = append(errosEmp, "WarehouseID invalid")
		}
	}

	if len(errosEmp) > 0 {
		err = errors.New(fmt.Sprintf("validation errors: %v", errosEmp))
	}
	return
}
