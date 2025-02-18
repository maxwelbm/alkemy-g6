package sellersctl

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

type SellerUpdateJSON struct {
	CID         *string `json:"cid,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
	Telephone   *string `json:"telephone,omitempty"`
	LocalityID  *int    `json:"locality_id,omitempty"`
}

func (j *SellerUpdateJSON) validate() (err error) {
	var validationErrors []string

	// Validate CID: must be positive if provided
	if j.CID != nil && *j.CID == "" {
		validationErrors = append(validationErrors, "error: attribute CID cannot be empty")
	}

	// Validate CompanyName: cannot be empty if provided
	if j.CompanyName != nil && *j.CompanyName == "" {
		validationErrors = append(validationErrors, "error: attribute CompanyName cannot be empty")
	}

	// Validate Address: cannot be empty if provided
	if j.Address != nil && *j.Address == "" {
		validationErrors = append(validationErrors, "error: attribute Address cannot be empty")
	}

	// Validate Telephone: cannot be empty if provided
	if j.Telephone != nil && *j.Telephone == "" {
		validationErrors = append(validationErrors, "error: attribute Telephone cannot be empty")
	}

	// Validate LocalityID: must be positive if provided
	if j.LocalityID != nil && *j.LocalityID < 1 {
		validationErrors = append(validationErrors, "error: attribute LocalityID must be positive")
	}

	// If there are validation errors, create an error with the details
	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	return err
}

// Update handles the HTTP PUT request to update a seller by ID.
//
// @Summary Update a seller
// @Description This endpoint updates a seller based on the provided ID in the URL parameter and the JSON request body.
// @Tags sellers
// @Accept json
// @Produce json
// @Param id path int true "Seller ID"
// @Param seller body SellerUpdateJSON true "Seller update data"
// @Success 200 {object} SellerResJSON "Success - The seller was successfully updated"
// @Failure 400 {object} response.ErrorResponse "Bad Request - The request ID is invalid, less than 1, or the request body is invalid"
// @Failure 400 {object} response.ErrorResponse "Bad Request - The seller cannot be updated due to a MySQL cannot add or update child row error"
// @Failure 409 {object} response.ErrorResponse "Conflict - The seller cannot be updated due to a MySQL duplicate entry error"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error - An unexpected error occurred during the update process"
// @Router /api/v1/sellers/{id} [patch]
func (controller *SellersDefault) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	var sellerRequest SellerUpdateJSON
	if err := json.NewDecoder(r.Body).Decode(&sellerRequest); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = sellerRequest.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	sellerToUpdate := createSellerDTO(sellerRequest)

	sellerUpdated, err := controller.sv.Update(id, sellerToUpdate)
	if err != nil {
		handleUpdateError(w, err)
		return
	}

	data := FullSellerJSON{
		ID:          sellerUpdated.ID,
		CID:         sellerUpdated.CID,
		CompanyName: sellerUpdated.CompanyName,
		Address:     sellerUpdated.Address,
		Telephone:   sellerUpdated.Telephone,
		LocalityID:  sellerUpdated.LocalityID,
	}
	res := SellerResJSON{
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}

func createSellerDTO(sellerRequest SellerUpdateJSON) models.SellerDTO {
	sellerToUpdate := models.SellerDTO{}
	if sellerRequest.CID != nil {
		sellerToUpdate.CID = *sellerRequest.CID
	}

	if sellerRequest.CompanyName != nil {
		sellerToUpdate.CompanyName = *sellerRequest.CompanyName
	}

	if sellerRequest.Address != nil {
		sellerToUpdate.Address = *sellerRequest.Address
	}

	if sellerRequest.Telephone != nil {
		sellerToUpdate.Telephone = *sellerRequest.Telephone
	}

	if sellerRequest.LocalityID != nil {
		sellerToUpdate.LocalityID = *sellerRequest.LocalityID
	}

	return sellerToUpdate
}

func handleUpdateError(w http.ResponseWriter, err error) {
	if errors.Is(err, models.ErrSellerNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == mysqlerr.CodeDuplicateEntry || mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
	}

	response.Error(w, http.StatusInternalServerError, err.Error())
}
