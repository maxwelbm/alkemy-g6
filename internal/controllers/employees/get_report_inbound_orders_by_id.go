package employeesctl

import (
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetReportInboundOrdersByID handles the HTTP request to retrieve report inbound orders for a specific employee.
// @Summary Retrieve report of inbound orders for an employee by their ID
// @Description Get report of inbound orders for an employee by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id query int true "Employee ID"
// @Success 200 {object} EmployeesResJSON "OK"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Employee not found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/employees/reportInboundOrders [get]
func (c *EmployeesController) GetReportInboundOrdersByID(w http.ResponseWriter, r *http.Request) {
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

	employees, err := c.SV.GetReportInboundOrdersByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if len(employees) == 0 {
		response.Error(w, http.StatusNotFound, "Employees not found")
		return
	}

	res := EmployeesResJSON{Data: employees}

	response.JSON(w, http.StatusOK, res)
}
