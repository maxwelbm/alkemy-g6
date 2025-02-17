package sellersctl

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetByID handles the HTTP GET request to retrieve a seller by ID.
//
// @Summary Get a seller by ID
// @Description This endpoint retrieves a seller based on the provided ID in the URL parameter.
// @Tags sellers
// @Produce json
// @Param id path int true "Seller ID"
// @Success 200 {object} SellerResJSON "Success - The seller was successfully retrieved"
// @Failure 400 {object} response.ErrorResponse "Bad Request - The request ID is invalid or less than 1"
// @Failure 404 {object} response.ErrorResponse "Not Found - The seller with the specified ID does not exist"
// @Router /api/v1/sellers/{id} [get]
func (controller *SellersDefault) GetByID(w http.ResponseWriter, r *http.Request) {
	// Extract the "id" parameter from the URL and convert it to an integer
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		// If conversion fails or id is less than 1, return a bad request error
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}

	// Call the service layer to get the seller by id
	seller, err := controller.sv.GetByID(id)
	if err != nil {
		// If the seller is not found, return a not found error
		response.Error(w, http.StatusNotFound, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

		return
	}

	// Create a FullSellerJSON struct with the seller data
	var data = FullSellerJSON{
		ID:          seller.ID,
		CID:         seller.CID,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
		LocalityID:  seller.LocalityID,
	}

	// Create a response struct with a success message and the seller data
	res := SellerResJSON{Data: data}

	// Send the JSON response with status OK
	response.JSON(w, http.StatusOK, res)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusOK, res))
}
