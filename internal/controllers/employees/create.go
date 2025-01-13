package employees_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
  models "github.com/maxwelbm/alkemy-g6/internal/models"

	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *EmployeesController) Create(w http.ResponseWriter, r *http.Request) {
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
		CardNumberID: employeesJson.CardNumberID,
		FirstName:    employeesJson.FirstName,
		LastName:     employeesJson.LastName,
		WarehouseID:  employeesJson.WarehouseID,
	}

	emp, err := c.SV.Create(employees)
	if err != nil {
		if errors.Is(err, service.ErrWareHousesServiceNotFound) {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())
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

func validateNewEmployees(employees EmployeesReqJSON) (err error) {

	var errosEmp []string
	if employees.CardNumberID == nil || *employees.CardNumberID == "" {
		errosEmp = append(errosEmp, "error: attribute CardNumberID cannot be empty")
	}

	if employees.FirstName == nil || *employees.FirstName == "" {
		errosEmp = append(errosEmp, "error: attribute FirstName cannot be empty")
	}

	if employees.LastName == nil || *employees.LastName == "" {
		errosEmp = append(errosEmp, "error: attribute LastName cannot be empty")
	}

	if employees.WarehouseID == nil {
		errosEmp = append(errosEmp, "error: attribute WarehouseID cannot be empty")
	} else if employees.WarehouseID == nil && *employees.WarehouseID <= 0 {
		errosEmp = append(errosEmp, "error: attribute WarehouseID invalid")
	}

	if len(errosEmp) > 0 {
		err = errors.New(fmt.Sprintf("validation errors: %v", errosEmp))
	}
	return
}
