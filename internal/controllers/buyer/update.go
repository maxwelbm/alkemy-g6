package buyers_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type BuyerUpdateJSON struct {
	CardNumberId *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}

func (j *BuyerUpdateJSON) validate() (err error) {
	var validationErrors []string

	// Validate FirstName: cannot be empty if provided
	if j.FirstName != nil && *j.FirstName == "" {
		validationErrors = append(validationErrors, "error: attribute FirstName cannot be empty")
	}

	// Validate LastName: cannot be empty if provided
	if j.LastName != nil && *j.LastName == "" {
		validationErrors = append(validationErrors, "error: attribute LastName cannot be empty")
	}

	// If there are validation errors, create an error with the details
	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	return
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
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 409 {object} ErrorResponse "Conflict"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /buyers/{id} [put]
func (ct *BuyersDefault) Update(w http.ResponseWriter, r *http.Request) {
	// Parse the buyer ID from the URL parameter and validate it
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Decode the JSON request body into a BuyerUpdateJSON struct
	var buyerRequest BuyerUpdateJSON
	if err := json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	// Validate the decoded request data
	if err := buyerRequest.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Create a BuyerDTO from the request data to update the buyer
	buyerToUpdate := models.BuyerDTO{
		CardNumberId: buyerRequest.CardNumberId,
		FirstName:    buyerRequest.FirstName,
		LastName:     buyerRequest.LastName,
	}

	// Call the service layer to update the buyer information
	buyerReturn, err := ct.SV.Update(id, buyerToUpdate)

	// Handle the case where no changes were made
	if errors.Is(err, models.ErrorNoChangesMade) {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		//  Handle the case where the buyer ID is not found
		if errors.Is(err, models.ErrBuyerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		// Handle the case where no changes were made
		if errors.Is(err, models.ErrorNoChangesMade) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		// Handle the case where a MySQL duplicate entry error occurred
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, respond with a 500 Internal Server Error status
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response data
	data := FullBuyerJSON{
		Id:           buyerReturn.Id,
		CardNumberId: buyerReturn.CardNumberId,
		FirstName:    buyerReturn.FirstName,
		LastName:     buyerReturn.LastName,
	}

	// Create the response JSON with a success message
	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}

	// Send the response back to the client
	response.JSON(w, http.StatusOK, res)
}
