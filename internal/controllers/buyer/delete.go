package buyersctl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Delete handles the HTTP DELETE request to remove a buyer by ID.
//
// @Summary Delete a buyer
// @Description Deletes a buyer from the system by their ID.
// @Tags buyers
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 204 {object} nil "No Content"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Router /api/v1/buyers/{id} [delete]
func (ct *BuyersDefault) Delete(w http.ResponseWriter, r *http.Request) {
	// Parse the buyer ID from the URL parameter and convert it to an integer.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		// If the ID is not valid, return a 400 Bad Request response.
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Attempt to delete the buyer by ID.
	err = ct.sv.Delete(id)
	if err != nil {
		// If the buyer ID is not found, return a 404 Not Found response.
		if errors.Is(err, models.ErrBuyerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		// For any other errors, return a 500 Internal Server Error response.
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// If the deletion is successful, return a 204 No Content response.
	response.JSON(w, http.StatusNoContent, nil)
}
