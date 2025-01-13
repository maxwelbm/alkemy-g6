package warehouses_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
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
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.Service.GetById(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	err = c.Service.Delete(id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotDeleteOrUpdateParentRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	res := WarehouseResJSON{
		Message: "Success",
		Data:    nil,
	}
	response.JSON(w, http.StatusNoContent, res)
}
