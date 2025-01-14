package productrecordsctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type ProductRecordCreateJSON struct {
	LastUpdateDate *string  `json:"last_update_date,omitempty"`
	PurchasePrice  *float64 `json:"purchase_price,omitempty"`
	SalePrice      *float64 `json:"sale_price,omitempty"`
	ProductID      *int     `json:"product_id,omitempty"`
}

func (j *ProductRecordCreateJSON) validate() (err error) {
	// Initialize a slice to hold validation error messages
	var validationErrors []string

	// Check if LastUpdateDate is nil and add an error message if it is
	if j.LastUpdateDate == nil {
		validationErrors = append(validationErrors, "error: last_update_date is required")
	}
	// Check if PurchasePrice is nil and add an error message if it is
	if j.PurchasePrice == nil {
		validationErrors = append(validationErrors, "error: purchase_price is required")
	}
	// Check if SalePrice is nil and add an error message if it is
	if j.SalePrice == nil {
		validationErrors = append(validationErrors, "error: sale_price is required")
	}
	// Check if ProductID is nil and add an error message if it is
	if j.ProductID == nil {
		validationErrors = append(validationErrors, "error: product_id is required")
	}
	// If there are any validation errors, create an error with all messages
	if len(validationErrors) > 0 {
		err = fmt.Errorf("validation errors: %v", validationErrors)
	}

	// Return the error (if any)
	return
}

// Create handles the creation of a new product record.
// @Summary Create a new product record
// @Description Create a new product record with the provided JSON data
// @Tags product_records
// @Accept json
// @Produce json
// @Param productRecord body ProductRecordCreateJSON true "Product Record Create JSON"
// @Success 201 {object} ProductRecordResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Error ao decodificar JSON"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/product_records [post]
func (controller *ProductRecordsDefault) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into productRecordRequest
	var productRecordRequest ProductRecordCreateJSON
	if err := json.NewDecoder(r.Body).Decode(&productRecordRequest); err != nil {
		// If there's an error decoding the JSON, respond with a bad request status
		response.Error(w, http.StatusBadRequest, "Error ao decodificar JSON")
		return
	}

	// Validate the request data
	if err := productRecordRequest.validate(); err != nil {
		// If validation fails, respond with a bad request status
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Map the request data to a ProductRecordDTO model
	productRecordToCreate := models.ProductRecordDTO{
		LastUpdateDate: *productRecordRequest.LastUpdateDate,
		PurchasePrice:  *productRecordRequest.PurchasePrice,
		SalePrice:      *productRecordRequest.SalePrice,
		ProductID:      *productRecordRequest.ProductID,
	}

	// Call the service layer to create the productRecord
	productRecordCreated, err := controller.sv.Create(productRecordToCreate)
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
	data := FullProductRecordJSON{
		ID:             productRecordCreated.ID,
		LastUpdateDate: productRecordCreated.LastUpdateDate,
		PurchasePrice:  productRecordCreated.PurchasePrice,
		SalePrice:      productRecordCreated.SalePrice,
		ProductID:      productRecordCreated.ProductID,
	}

	// Create the response JSON
	res := ProductRecordResJSON{
		Message: "Success",
		Data:    data,
	}

	// Respond with the created status and the response JSON
	response.JSON(w, http.StatusCreated, res)

}
