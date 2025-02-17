package buyersctl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type BuyerUpdateJSON struct {
	CardNumberID *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}

func (j *BuyerUpdateJSON) validate() (err error) {
	var validationErrors []string

	if j.CardNumberID != nil && *j.CardNumberID == "" {
		validationErrors = append(validationErrors, "error: attribute CardNumberID cannot be empty")
	}

	if j.FirstName != nil && *j.FirstName == "" {
		validationErrors = append(validationErrors, "error: attribute FirstName cannot be empty")
	}

	if j.LastName != nil && *j.LastName == "" {
		validationErrors = append(validationErrors, "error: attribute LastName cannot be empty")
	}

	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	return err
}

// Update handles the HTTP request to update a buyer's information.
// @Summary Update buyer
// @Description Update an existing buyer's details by ID
// @Tags buyers
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Param buyer body BuyerUpdateJSON true "Buyer Update JSON"
// @Success 200 {object} BuyerResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/buyers/{id} [patch]
func (ct *BuyersDefault) Update(w http.ResponseWriter, r *http.Request) {
	// Parse the buyer ID from the URL parameter and validate it
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}

	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))

		return
	}

	// Decode the JSON request body into a BuyerUpdateJSON struct
	var buyerRequest BuyerUpdateJSON
	if err = json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}

	// Validate the decoded request data
	if err = buyerRequest.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusUnprocessableEntity, err.Error()))

		return
	}

	// Create a BuyerDTO from the request data to update the buyer
	buyerToUpdate := models.BuyerDTO{
		CardNumberID: buyerRequest.CardNumberID,
		FirstName:    buyerRequest.FirstName,
		LastName:     buyerRequest.LastName,
	}

	// Call the service layer to update the buyer information
	buyerReturn, err := ct.sv.Update(id, buyerToUpdate)

	if err != nil {
		// Handle the case where no changes were made
		if errors.Is(err, models.ErrorNoChangesMade) {
			response.Error(w, http.StatusBadRequest, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

			return
		}
		//  Handle the case where the buyer ID is not found
		if errors.Is(err, models.ErrBuyerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

			return
		}
		// Handle the case where a MySQL duplicate entry error occurred
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusConflict, err.Error()))

			return
		}
		// For any other error, respond with a 500 Internal Server Error status
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	// Prepare the response data
	data := FullBuyerJSON{
		ID:           buyerReturn.ID,
		CardNumberID: buyerReturn.CardNumberID,
		FirstName:    buyerReturn.FirstName,
		LastName:     buyerReturn.LastName,
	}

	// Create the response JSON with a success message
	res := BuyerResJSON{
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}

	// Send the response back to the client
	response.JSON(w, http.StatusOK, res)
}
