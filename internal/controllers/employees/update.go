package employeesctl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Update handles the HTTP request to update an employee by their ID.
//
// It extracts the employee ID from the URL parameters, validates it, and
// fetches the employee details from the service layer. If the ID is invalid
// or the employee is not found, it returns an appropriate error response.
//
// @Summary Update employee by ID
// @Description Update an existing employee by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body EmployeesReqJSON true "Employee JSON"
// @Success 200 {object} EmployeesResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "employee not found"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Router /api/v1/employees/{id} [patch]
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

	newEmployees := models.EmployeeDTO{}

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

	emp, err := c.sv.Update(newEmployees, id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case mysqlerr.CodeDuplicateEntry:
				response.Error(w, http.StatusConflict, err.Error())
				return
			case mysqlerr.CodeCannotAddOrUpdateChildRow:
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	data := EmployeesResJSON{
		Message: http.StatusText(http.StatusOK),
		Data: EmployeesFullJSON{
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
		err = fmt.Errorf("validation errors: %v", errosEmp)
	}

	return err
}
