package employees_controller

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetAll handles the HTTP GET request to retrieve all employees.
// It fetches employee data from the service layer, converts it into a list
// of EmployeesAttributes, and sends a JSON response with the data.
//
// @Summary Retrieve all employees
// @Description Fetches all employees from the database and returns them as JSON.
// @Tags employees
// @Produce json
// @Success 200 {object} EmployeesFinalJSON "OK - The employees were successfully retrieved"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error - An unexpected error occurred during the retrieval process"
// @Router /api/v1/employees [get]
func (c *EmployeesController) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := c.SV.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var list []EmployeesAttributes
	for _, value := range employees {
		list = append(list, EmployeesAttributes{
			ID:           value.ID,
			CardNumberID: value.CardNumberID,
			FirstName:    value.FirstName,
			LastName:     value.LastName,
			WarehouseID:  value.WarehouseID,
		})
	}

	responseEmp := EmployeesFinalJSON{
		Data: list,
	}

	response.JSON(w, http.StatusOK, responseEmp)
}
