package buyersctl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type BuyerCreateJSON struct {
	CardNumberID *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}

func (j *BuyerCreateJSON) validate() (err error) {
	var validationErrors, nilPointerErrors []string

	if j.FirstName == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute FirstName cannot be nil")
	} else if j.FirstName != nil && *j.FirstName == "" {
		validationErrors = append(validationErrors, "error: attribute FirstName cannot be empty")
	}

	if j.LastName == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute LastName cannot be nil")
	} else if j.LastName != nil && *j.LastName == "" {
		validationErrors = append(validationErrors, "error: attribute LastName cannot be empty")
	}

	if j.CardNumberID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CardNumberID cannot be nil")
	} else if j.CardNumberID != nil && *j.CardNumberID == "" {
		validationErrors = append(validationErrors, "error: attribute CardNumberID cannot be empty")
	}

	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, nilPointerErrors...)
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}

	return err
}

// Create handles the creation of a new buyer.
// @Summary Create a new buyer
// @Description Create a new buyer with the provided details
// @Tags buyers
// @Accept json
// @Produce json
// @Param buyer body BuyerCreateJSON true "Buyer details"
// @Success 201 {object} BuyerResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/buyers [post]
func (ct *BuyersDefault) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a BuyerCreateJSON struct
	var buyerRequest BuyerCreateJSON
	if err := json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		// If there's an error decoding the request, respond with a 400 Bad Request status
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the decoded request data
	if err := buyerRequest.validate(); err != nil {
		// If validation fails, respond with a 400 Bad Request status
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Create a BuyerDTO from the request data
	buyerToCreate := models.BuyerDTO{
		CardNumberID: buyerRequest.CardNumberID,
		FirstName:    buyerRequest.FirstName,
		LastName:     buyerRequest.LastName,
	}

	// Attempt to create the buyer using the service
	buyerCreated, err := ct.sv.Create(buyerToCreate)
	if err != nil {
		// Handle specific MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			// If there's a duplicate entry, respond with a 409 Conflict status
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For other errors, respond with a 500 Internal Server Error status
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	// Prepare the response data
	data := FullBuyerJSON{
		ID:           buyerCreated.ID,
		CardNumberID: buyerCreated.CardNumberID,
		FirstName:    buyerCreated.FirstName,
		LastName:     buyerCreated.LastName,
	}

	// Create the response JSON
	res := BuyerResJSON{
		Message: http.StatusText(http.StatusCreated),
		Data:    data,
	}

	// Respond with a 201 Created status and the response JSON
	response.JSON(w, http.StatusCreated, res)
}
