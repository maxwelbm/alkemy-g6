package buyersctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type BuyerCreateJson struct {
	CardNumberId *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}

func (j *BuyerCreateJson) validate() (err error) {
	var validationErrors []string

	if j.FirstName == nil {
		validationErrors = append(validationErrors, "error: first_name is required")
	}
	if j.LastName == nil {
		validationErrors = append(validationErrors, "error: last_name is required")
	}
	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	return
}

// Create handles the creation of a new buyer.
// @Summary Create a new buyer
// @Description Create a new buyer with the provided details
// @Tags buyers
// @Accept json
// @Produce json
// @Param buyer body BuyerCreateJson true "Buyer details"
// @Success 201 {object} BuyerResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/buyers [post]
func (ct *BuyersDefault) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a BuyerCreateJson struct
	var buyerRequest BuyerCreateJson
	if err := json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		// If there's an error decoding the request, respond with a 400 Bad Request status
		response.JSON(w, http.StatusBadRequest, err.Error())
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
		CardNumberId: buyerRequest.CardNumberId,
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
		Id:           buyerCreated.Id,
		CardNumberId: buyerCreated.CardNumberId,
		FirstName:    buyerCreated.FirstName,
		LastName:     buyerCreated.LastName,
	}

	// Create the response JSON
	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}

	// Respond with a 201 Created status and the response JSON
	response.JSON(w, http.StatusCreated, res)
}
