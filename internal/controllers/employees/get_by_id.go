package employeesctl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetByID handles the HTTP request to retrieve an employee by their ID.
// It extracts the employee ID from the URL parameters, validates it, and
// fetches the employee details from the service layer. If the ID is invalid
// or the employee is not found, it returns an appropriate error response.
//
// @Summary Retrieve employee by ID
// @Description Get employee details by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} EmployeesResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "employee not found"
// @Router /api/v1/employees/{id} [get]
func (c *EmployeesController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if id <= 0 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	employees, err := c.sv.GetByID(id)
	if err != nil {
		if errors.Is(err, models.ErrEmployeeNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	data := EmployeesFullJSON{
		ID:           employees.ID,
		CardNumberID: employees.CardNumberID,
		FirstName:    employees.FirstName,
		LastName:     employees.LastName,
		WarehouseID:  employees.WarehouseID,
	}

	res := EmployeesResJSON{Data: data}

	response.JSON(w, http.StatusOK, res)
}
