package sellersctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type SellerCreateJSON struct {
	CID         *string `json:"cid,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
	Telephone   *string `json:"telephone,omitempty"`
	LocalityID  *int    `json:"locality_id,omitempty"`
}

func (j *SellerCreateJSON) validate() (err error) {
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
	// Check if Telephone is nil and add an error message if it is
	if j.Telephone == nil {
		validationErrors = append(validationErrors, "error: telephone is required")
	}
	// Check if LocalityID is nil and add an error message if it is
	if j.LocalityID == nil {
		validationErrors = append(validationErrors, "error: locality_id is required")
	}
	// If there are any validation errors, create an error with all messages
	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	// Return the error (if any)
	return
}

// Create handles the creation of a new seller.
//
// @Summary Create a new seller
// @Description This endpoint creates a new seller based on the provided JSON request body.
// @Tags sellers
// @Accept json
// @Produce json
// @Param seller body SellerCreateJSON true "Seller Create JSON"
// @Success 201 {object} SellerResJSON "Seller created successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid request data or JSON decoding error"
// @Failure 409 {object} response.ErrorResponse "Conflict - Duplicate entry"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/sellers [post]
func (controller *SellersDefault) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into sellerRequest
	var sellerRequest SellerCreateJSON
	if err := json.NewDecoder(r.Body).Decode(&sellerRequest); err != nil {
		// If there's an error decoding the JSON, respond with a bad request status
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the request data
	if err := sellerRequest.validate(); err != nil {
		// If validation fails, respond with a bad request status
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Map the request data to a SellerDTO model
	sellerToCreate := models.SellerDTO{
		CID:         *sellerRequest.CID,
		CompanyName: *sellerRequest.CompanyName,
		Address:     *sellerRequest.Address,
		Telephone:   *sellerRequest.Telephone,
		LocalityID:  *sellerRequest.LocalityID,
	}

	// Call the service layer to create the seller
	sellerCreated, err := controller.sv.Create(sellerToCreate)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response data
	data := FullSellerJSON{
		ID:          sellerCreated.ID,
		CID:         sellerCreated.CID,
		CompanyName: sellerCreated.CompanyName,
		Address:     sellerCreated.Address,
		Telephone:   sellerCreated.Telephone,
		LocalityID:  sellerCreated.LocalityID,
	}

	// Create the response JSON
	res := SellerResJSON{
		Message: "Success",
		Data:    data,
	}

	// Respond with the created status and the response JSON
	response.JSON(w, http.StatusCreated, res)

}
