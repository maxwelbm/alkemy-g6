package employeesctl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetReportInboundOrders handles the HTTP request to retrieve report inbound orders for a specific employee.
// @Summary Retrieve report of inbound orders for an employee by their ID
// @Description Get report of inbound orders for an employee by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id query int true "Employee ID"
// @Success 200 {object} EmployeesResJSON "OK"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "employee not found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/employees/reportInboundOrders [get]
func (c *EmployeesController) GetReportInboundOrders(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	var id int

	var err error

	if idString != "" {
		id, err = strconv.Atoi(idString)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if id <= 0 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}

	reports, err := c.sv.GetReportInboundOrders(id)
	if err != nil {
		if errors.Is(err, models.ErrEmployeeNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	data := make([]ReportInboundOrdersFullJSON, 0, len(reports))
	for _, r := range reports {
		data = append(data,
			ReportInboundOrdersFullJSON{
				ID:           r.ID,
				CardNumberID: r.CardNumberID,
				FirstName:    r.FirstName,
				LastName:     r.LastName,
				WarehouseID:  r.WarehouseID,
				CountReports: r.CountReports,
			},
		)
	}

	res := EmployeesResJSON{Data: data}

	response.JSON(w, http.StatusOK, res)
}
