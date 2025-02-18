package warehousesctl

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

// Delete handles the deletion of a warehouse by its ID.
// @Summary Delete a warehouse
// @Description Deletes a warehouse by its ID from the database.
// @Tags warehouses
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 204 {object} WarehouseResJSON "No content"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 404 {object} response.ErrorResponse "Warehouse not found"
// @Failure 409 {object} response.ErrorResponse "Conflict - cannot delete or update parent row"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/warehouses/{id} [delete]
func (c *WarehouseDefault) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	// If the ID is less than 1, return a 400 Bad Request error
	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = c.sv.Delete(id)

	if err != nil {
		// Handle if section not found
		if errors.Is(err, models.ErrWareHouseNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		// Handle no changes made
		if errors.Is(err, models.ErrorNoChangesMade) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		// Handle MySQL conflict dependencies
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotDeleteOrUpdateParentRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// Handle other internal server errors
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
