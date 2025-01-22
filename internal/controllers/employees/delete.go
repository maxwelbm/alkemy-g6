package employeesctl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Delete handles the deletion of a employee by its ID.
// @Summary Delete a employee
// @Description Delete a employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 204
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/employees/{id} [delete]
func (c *EmployeesController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	
	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = c.sv.Delete(id)

	if err != nil {
		// Handle if section not found
		if errors.Is(err, models.ErrEmployeeNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		// Handle MySQL conflict dependencies
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}	

	data := map[string]string{
		"Message": "Success delete",
	}

	response.JSON(w, http.StatusNoContent, data)
}
