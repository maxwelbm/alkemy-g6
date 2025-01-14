package sellersctl

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

// Delete handles the HTTP DELETE request to remove a seller by ID.
//
// @Summary Delete a seller
// @Description This endpoint deletes a seller based on the provided ID in the URL parameter.
// @Tags sellers
// @Produce json
// @Param id path int true "Seller ID"
// @Success 204 "No Content - The seller was successfully deleted"
// @Failure 400 {object} response.ErrorResponse "Bad Request - The request ID is invalid or less than 1"
// @Failure 404 {object} response.ErrorResponse "Not Found - The seller with the specified ID does not exist"
// @Failure 409 {object} response.ErrorResponse "Conflict - The seller cannot be deleted due to a MySQL foreign key constraint error"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error - An unexpected error occurred during the deletion process"
// @Router /api/v1/sellers/{id} [delete]
func (controller *SellersDefault) Delete(w http.ResponseWriter, r *http.Request) {
	// Convert the URL parameter "id" to an integer
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		// If conversion fails or id is less than 1, return a bad request error
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Attempt to delete the seller by id
	err = controller.sv.Delete(id)
	if err != nil {
		// If the seller is not found, return a not found error
		if errors.Is(err, models.ErrSellerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		// Check if the error is a MySQL foreign key constraint error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotDeleteOrUpdateParentRow {
			// If it is, return a conflict error
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, return an internal server error
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	// If deletion is successful, return no content status
	response.JSON(w, http.StatusNoContent, nil)
}
