package employees_controller

import (
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetReportInboundOrdersById handles the HTTP request to retrieve a report of inbound orders by employee ID.
// It expects an "id" query parameter in the request URL. If the "id" parameter is not provided or is invalid,
// it returns a 400 Bad Request response. If no employees are found for the given ID, it returns a 404 Not Found response.
// On success, it returns a JSON response with the list of employees and a 200 OK status.
//
// @Summary Retrieve employee by ID
// @Description Get employee details by their ID
// @Tags employees
// @Accept json
// @Produce json
// @param w http.ResponseWriter - the response writer to send the HTTP response
// @param r *http.Request - the HTTP request containing the query parameter "id"
// @Success 200 {object} EmployeesResJSON "Success"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 404 {object} ErrorResponse "Employee not found"
// @Router /api/v1/employees/reportInboundOrders?id={} [get]

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

	if len(employees) == 0 {
		response.Error(w, http.StatusNotFound, "Employees not found")
		return
	}

	res := EmployeesResJSON{
		Message: "Success",
		Data:    employees,
	}

	response.JSON(w, http.StatusOK, res)
}
