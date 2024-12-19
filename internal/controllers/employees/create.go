package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *Employees) Create(w http.ResponseWriter, r *http.Request) {
	var employeesJson EmployeesReqJSON
	err := json.NewDecoder(r.Body).Decode(&employeesJson)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, nil)
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
