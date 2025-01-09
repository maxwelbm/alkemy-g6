package buyers_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetById handles the HTTP request to retrieve a buyer by their ID.
// @Summary Get buyer by ID
// @Description Get details of a buyer by their ID
// @Tags buyers
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} BuyerResJSON "Success"
// @Failure 400 {object} ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} ErrorResponse "Buyer not found"
// @Router /buyers/{id} [get]
func (ct *BuyersDefault) GetById(w http.ResponseWriter, r *http.Request) {
	// Parse the buyer ID from the URL parameter and convert it to an integer
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		// If the ID is invalid, return a 400 Bad Request error
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Retrieve the buyer details from the service layer using the ID
	buyer, err := ct.SV.GetById(id)
	if err != nil {
		// If the buyer is not found, return a 404 Not Found error
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	// Map the buyer details to the response JSON structure
	var data FullBuyerJSON
	data = FullBuyerJSON{
		Id:           buyer.Id,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}

	// Create the response JSON with a success message and the buyer data
	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}

	// Send the response JSON with a 200 OK status
	response.JSON(w, http.StatusOK, res)
}
