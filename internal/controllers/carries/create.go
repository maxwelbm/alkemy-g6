package carriesctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type CarriesCreateJSON struct {
	CID         *string `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phone_number"`
	LocalityID  *int    `json:"locality_id"`
}

func (j *CarriesCreateJSON) validate() (err error) {
	// Initialize a slice to hold validation error messages
	var validationErrors []string

	// Check if CID is nil and add an error message if it is
	if j.CID == nil {
		validationErrors = append(validationErrors, "error: cid is required")
	}
	// Check if CompanyName is nil and add an error message if it is
	if j.CompanyName == nil {
		validationErrors = append(validationErrors, "error: company_name is required")
	}
	// Check if Address is nil and add an error message if it is
	if j.Address == nil {
		validationErrors = append(validationErrors, "error: address is required")
	}
	// Check if PhoneNumber is nil and add an error message if it is
	if j.PhoneNumber == nil {
		validationErrors = append(validationErrors, "error: phone_number is required")
	}
	// Check if LocalityID is nil and add an error message if it is
	if j.LocalityID == nil {
		validationErrors = append(validationErrors, "error: locality_id is required")
	}
	// If there are any validation errors, create an error with all messages
	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	return
}

// Create creates a new carry
// @Summary Create a new carry
// @Description Create a new carry in the database
// @Tags carries
// @Produce json
// @Param carry body CarriesCreateJSON true "Carry to create"
// @Success 201 {object} CarriesCreateJSON "OK"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/carries [post]
func (ctl *CarriesDefault) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into sellerRequest
	var carryRequest CarriesCreateJSON
	if err := json.NewDecoder(r.Body).Decode(&carryRequest); err != nil {
		// If there's an error decoding the JSON, respond with a bad request status
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the request data
	if err := carryRequest.validate(); err != nil {
		// If validation fails, respond with a bad request status
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Map the request data to a CarriesDTO model
	carryToCreate := models.CarryDTO{
		CID:         *carryRequest.CID,
		CompanyName: *carryRequest.CompanyName,
		Address:     *carryRequest.Address,
		PhoneNumber: *carryRequest.PhoneNumber,
		LocalityID:  *carryRequest.LocalityID,
	}

	// Call the service layer to create the carry
	carryCreated, err := ctl.sv.Create(carryToCreate)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error or cannot add or update child row error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok &&
			mysqlErr.Number == mysqlerr.CodeDuplicateEntry ||
			mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}

		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	// Prepare the response data
	data := FullCarryJSON{
		ID:          carryCreated.ID,
		CID:         carryCreated.CID,
		CompanyName: carryCreated.CompanyName,
		Address:     carryCreated.Address,
		PhoneNumber: carryCreated.PhoneNumber,
		LocalityID:  carryCreated.LocalityID,
	}

	// Create the response JSON
	res := CarriesResJSON{
		Message: http.StatusText(http.StatusCreated),
		Data:    data,
	}
	response.JSON(w, http.StatusCreated, res)
}
