package buyersctl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetByID handles the HTTP request to retrieve a buyer by their ID.
// @Summary Get buyer by ID
// @Description Get details of a buyer by their ID
// @Tags buyers
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} BuyerResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} response.ErrorResponse "Buyer not found"
// @Router /api/v1/buyers/{id} [get]
func (ct *BuyersDefault) GetByID(w http.ResponseWriter, r *http.Request) {
	// Parse the buyer ID from the URL parameter and validate it
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	// id, err := strconv.Atoi(r.URL.Path[len("/api/v1/buyers/"):])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Retrieve the buyer details from the service layer using the ID
	buyer, err := ct.sv.GetByID(id)
	if err != nil {
		// If the buyer ID is not found, return a 404 Not Found response
		if errors.Is(err, models.ErrBuyerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		// For any other errors, return a 500 Internal Server Error response
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	// Map the buyer details to the response JSON structure
	var data FullBuyerJSON
	data.ID = buyer.ID
	data.CardNumberID = buyer.CardNumberID
	data.FirstName = buyer.FirstName
	data.LastName = buyer.LastName

	// Create the response JSON with a success message and the buyer data
	res := BuyerResJSON{Data: data}

	// Send the response JSON with a 200 OK status
	response.JSON(w, http.StatusOK, res)
}
