package buyers_controller

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetAll handles the HTTP request to retrieve all buyers.
// @Summary Get all buyers
// @Description Retrieve a list of all buyers from the database
// @Tags buyers
// @Produce application/json
// @Success 200 {object} BuyerResJSON "Success"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/buyers [get]
func (ct *BuyersDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	// Retrieve all buyers from the service layer
	buyers, err := ct.sv.GetAll()

	// Check for errors in retrieving buyers
	if err != nil {
		// Send an internal server error response if there is an error
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Initialize a slice to hold the response data
	var data []FullBuyerJSON
	// Iterate over the retrieved buyers and map them to the response format
	for _, value := range buyers {
		new := FullBuyerJSON{
			Id:           value.Id,
			CardNumberId: value.CardNumberId,
			FirstName:    value.FirstName,
			LastName:     value.LastName,
		}

		// Append the mapped buyer data to the response slice
		data = append(data, new)
	}

	// Create the response JSON structure
	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}

	// Send the JSON response with status OK
	response.JSON(w, http.StatusOK, res)
}
