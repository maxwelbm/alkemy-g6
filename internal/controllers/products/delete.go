package productsctl

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

// Delete handles the deletion of a product by its ID.
// @Summary Delete a product
// @Description Deletes a product by its ID from the database.
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 {object} ProductResJSON "No content"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 404 {object} response.ErrorResponse "Product not found"
// @Failure 409 {object} response.ErrorResponse "Conflict - cannot delete or update parent row"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id} [delete]
func (p *ProductsDefault) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = p.SV.Delete(id)
	if errors.Is(err, models.ErrProductNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotDeleteOrUpdateParentRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	res := ProductResJSON{Message: "No content"}
	response.JSON(w, http.StatusNoContent, res)
}
