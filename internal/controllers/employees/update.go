package employees_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *EmployeesController) Update(w http.ResponseWriter, r *http.Request) {
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

	err = validateUpdateEmployees(employeesJSON)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	newEmployees := models.EmployeesDTO{}

	newEmployees.ID = employeesJSON.ID
	if employeesJSON.CardNumberID != nil {
		newEmployees.CardNumberID = employeesJSON.CardNumberID
	}

	if employeesJSON.FirstName != nil {
		newEmployees.FirstName = employeesJSON.FirstName
	}

	if employeesJSON.LastName != nil {
		newEmployees.LastName = employeesJSON.LastName
	}

	if employeesJSON.WarehouseID != nil {
		newEmployees.WarehouseID = employeesJSON.WarehouseID
	}

	emp, err := c.SV.Update(newEmployees, id)
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
	con := false
	if e.CardNumberID != nil && *e.CardNumberID == "" {
		errosEmp = append(errosEmp, "error: attribute CardNumberID cannot be empty")
		con = true
	}

	// FirstName
	if e.FirstName != nil && *e.FirstName == "" {
		if !con {
			errosEmp = append(errosEmp, "error: attribute FirstName cannot be empty")
		} else {
			errosEmp = append(errosEmp, "- error: attribute FirstName cannot be empty")
		}
		con = true
	}

	// LastName
	if e.LastName != nil && *e.LastName == "" {
		if !con {
			errosEmp = append(errosEmp, "error: attribute LastName cannot be empty")
		} else {
			errosEmp = append(errosEmp, "- error: attribute LastName cannot be empty")
		}
		con = true
	}

	// WarehouseID
	if e.WarehouseID != nil && *e.WarehouseID <= 0 {
		if !con {
			errosEmp = append(errosEmp, "error: attribute WarehouseID must be positive")
		} else {
			errosEmp = append(errosEmp, "- error: attribute WarehouseID must be positive")
		}
	}

	if len(errosEmp) > 0 {
		err = errors.New(fmt.Sprintf("validation errors: %v", errosEmp))
	}
	return
}
