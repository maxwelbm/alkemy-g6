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

//nolint:gocyclo
func (j *CarriesCreateJSON) validate() (err error) {
	// Initialize a slice to hold validation error messages
	var validationErrors []string

	// Check if CID is nil and add an error message if it is
	if j.CID == nil {
		validationErrors = append(validationErrors, "error: cid cannot be nil")
	}
	// Check if CID is empty and add an error message if it is
	if j.CID != nil && *j.CID == "" {
		validationErrors = append(validationErrors, "error: cid cannot be empty")
	}
	// Check if CompanyName is nil and add an error message if it is
	if j.CompanyName == nil {
		validationErrors = append(validationErrors, "error: company_name cannot be nil")
	}
	// Check if CompanyName is empty and add an error message if it is
	if j.CompanyName != nil && *j.CompanyName == "" {
		validationErrors = append(validationErrors, "error: company_name cannot be empty")
	}
	// Check if Address is nil and add an error message if it is
	if j.Address == nil {
		validationErrors = append(validationErrors, "error: address cannot be nil")
	}
	// Check if Address is empty and add an error message if it is
	if j.Address != nil && *j.Address == "" {
		validationErrors = append(validationErrors, "error: address cannot be empty")
	}
	// Check if PhoneNumber is nil and add an error message if it is
	if j.PhoneNumber == nil {
		validationErrors = append(validationErrors, "error: phone_number cannot be nil")
	}
	// Check if PhoneNumber is empty and add an error message if it is
	if j.PhoneNumber != nil && *j.PhoneNumber == "" {
		validationErrors = append(validationErrors, "error: phone_number cannot be empty")
	}
	// Check if LocalityID is nil and add an error message if it is
	if j.LocalityID == nil {
		validationErrors = append(validationErrors, "error: locality_id cannot be nil")
	}
	// Check if LocalityID is 0 and add an error message if it is
	if j.LocalityID != nil && *j.LocalityID == 0 {
		validationErrors = append(validationErrors, "error: locality_id cannot be 0")
	}
	// If there are any validation errors, create an error with all messages
	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	return err
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
		CID:         carryRequest.CID,
		CompanyName: carryRequest.CompanyName,
		Address:     carryRequest.Address,
		PhoneNumber: carryRequest.PhoneNumber,
		LocalityID:  carryRequest.LocalityID,
	}

	// Call the service layer to create the carry
	carryCreated, err := ctl.sv.Create(carryToCreate)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error or cannot add or update child row error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok &&
			(mysqlErr.Number == mysqlerr.CodeDuplicateEntry ||
				mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow) {
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
