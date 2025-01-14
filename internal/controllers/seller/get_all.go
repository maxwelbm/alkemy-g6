package sellersctl

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetAll handles the HTTP GET request to retrieve all sellers.
//
// @Summary Retrieve all sellers
// @Description This endpoint retrieves all sellers from the database using the service layer.
// @Tags sellers
// @Produce json
// @Success 200 {object} SellerResJSON "OK - The sellers were successfully retrieved"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error - An unexpected error occurred during the retrieval process"
// @Router /api/v1/sellers [get]
func (controller *SellersDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	// Retrieve all sellers using the service layer
	sellers, err := controller.sv.GetAll()

	// Check for errors in retrieving sellers
	if err != nil {
		// Send an error response if retrieval fails
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Initialize a slice to hold the full seller JSON data
	data := make([]FullSellerJSON, len(sellers))

	for _, value := range sellers {
		// Create a new FullSellerJSON object for each seller
		seller := FullSellerJSON{
			ID:          value.ID,
			CID:         value.CID,
			CompanyName: value.CompanyName,
			Address:     value.Address,
			Telephone:   value.Telephone,
			LocalityID:  value.LocalityID,
		}

		// Append the seller object to the data slice
		data = append(data, seller)
	}

	// Create a response object with a success message and the data
	res := SellerResJSON{
		Message: "Success",
		Data:    data,
	}

	// Send the JSON response with status OK
	response.JSON(w, http.StatusOK, res)
}
