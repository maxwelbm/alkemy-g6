package employees_controller

import (
	"net/http"
	"strconv"

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
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 404 {object} ErrorResponse "Employee not found"
// @Router /employees/{id} [get]
func (c *EmployeesController) GetReportInboundOrdersById(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	var id int
	var err error

	if idString != "" {
		id, err = strconv.Atoi(idString)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid ID format")
			return
		}
	}

	employees, err := c.SV.GetReportInboundOrdersById(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	res := EmployeesResJSON{
		Message: "Sucess",
		Data:    employees,
	}

	response.JSON(w, http.StatusOK, res)
}
