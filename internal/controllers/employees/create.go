package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *Employees) Create(w http.ResponseWriter, r *http.Request) {
	var employeesJson EmployeesReqJSON
	err := json.NewDecoder(r.Body).Decode(&employeesJson)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Body invalid")
		return
	}

	err = validateNewEmployees(employeesJson)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	employees := models.EmployeesDTO{
		CardNumberID: *employeesJson.CardNumberID,
		FirstName:    *employeesJson.FirstName,
		LastName:     *employeesJson.LastName,
		WarehouseID:  *employeesJson.WarehouseID,
	}

	emp, err := c.sv.Create(employees)
	if errors.Is(err, repository.ErrEmployeesRepositoryDuplicatedCode) {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}

	if errors.Is(err, service.ErrWareHousesServiceNotFound) {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
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
	response.JSON(w, http.StatusCreated, data)
}

func validateNewEmployees(employees EmployeesReqJSON) error {

	if employees.CardNumberID == nil || *employees.CardNumberID == "" {
		return errors.New("CardNumberID cannot be empty")
	}

	if employees.FirstName == nil || *employees.FirstName == "" {
		return errors.New("FirstName cannot be empty")
	}

	if employees.LastName == nil || *employees.LastName == "" {
		return errors.New("LastName cannot be empty")
	}

	if employees.WarehouseID == nil {
		return errors.New("WarehouseID cannot be empty")
	}

	if *employees.WarehouseID <= 0 {
		return errors.New("WarehouseID invalid")
	}
	return nil
}
