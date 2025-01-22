package employeesctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	models "github.com/maxwelbm/alkemy-g6/internal/models"

	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Create handles the creation of a new employee.
// @Summary Create a new employee
// @Description Create a new employee with the provided JSON data
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body EmployeesReqJSON true "Employee JSON"
// @Success 201 {object} EmployeesResJSON "Created"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/employees [post]
func (c *EmployeesController) Create(w http.ResponseWriter, r *http.Request) {
	var employeesJSON EmployeesReqJSON
	err := json.NewDecoder(r.Body).Decode(&employeesJSON)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateNewEmployees(employeesJSON)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	employees := models.EmployeeDTO{
		CardNumberID: employeesJSON.CardNumberID,
		FirstName:    employeesJSON.FirstName,
		LastName:     employeesJSON.LastName,
		WarehouseID:  employeesJSON.WarehouseID,
	}

	emp, err := c.sv.Create(employees)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && (mysqlErr.Number == mysqlerr.CodeDuplicateEntry || mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow) {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	data := EmployeesResJSON{
		Message: "Success created",
		Data: EmployeeFullJSON{
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
	var validationErrors, nilPointerErrors []string

	if employees.CardNumberID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CardNumberID cannot be nil")
	} else if *employees.CardNumberID == "" {
		validationErrors = append(validationErrors, "error: attribute CardNumberID cannot be empty")
	}

	if employees.FirstName == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute FirstName cannot be nil")
	} else if *employees.FirstName == "" {
		validationErrors = append(validationErrors, "error: attribute FirstName cannot be empty")
	}

	if employees.LastName == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute LastName cannot be nil")
	} else if *employees.LastName == "" {
		validationErrors = append(validationErrors, "error: attribute LastName cannot be empty")
	}

	if employees.WarehouseID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute WarehouseID cannot be empty")
	} else if employees.WarehouseID != nil && *employees.WarehouseID <= 0 {
		validationErrors = append(validationErrors, "error: attribute WarehouseID invalid")
	}

	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, nilPointerErrors...)
		allErrors = append(allErrors, validationErrors...)

		err = fmt.Errorf("validation errors: %v", allErrors)
	}

	return err
}
