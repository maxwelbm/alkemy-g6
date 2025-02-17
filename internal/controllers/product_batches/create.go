package productbatchesctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

//nolint:gocyclo
func (prodBatch *NewProductBatchesReqJSON) validateCreate() (err error) {
	var validationErrors, nilPointerErrors []string

	if prodBatch.BatchNumber == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute BatchNumber cannot be nil")
	} else if *prodBatch.BatchNumber == "" {
		validationErrors = append(validationErrors, "error: attribute BatchNumber cannot be empty")
	}

	if prodBatch.InitialQuantity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute InitialQuantity cannot be nil")
	} else if *prodBatch.InitialQuantity <= 0 {
		validationErrors = append(validationErrors, "error: attribute InitialQuantity cannot be negative")
	}

	if prodBatch.CurrentQuantity == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentQuantity cannot be nil")
	} else if *prodBatch.CurrentQuantity <= 0 {
		validationErrors = append(validationErrors, "error: attribute CurrentQuantity cannot be negative")
	}

	if prodBatch.CurrentTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute CurrentTemperature cannot be nil")
	}

	if prodBatch.MinimumTemperature == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute MinimumTemperature cannot be nil")
	}

	if prodBatch.DueDate == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute DueDate cannot be nil")
	}

	if prodBatch.ManufacturingDate == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ManufacturingDate cannot be nil")
	}

	if prodBatch.ManufacturingHour == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ManufacturingHour cannot be nil")
	}

	if prodBatch.ProductID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ProductID cannot be nil")
	} else if *prodBatch.ProductID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductID must be positive")
	}

	if prodBatch.SectionID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute SectionID cannot be nil")
	} else if *prodBatch.SectionID <= 0 {
		validationErrors = append(validationErrors, "error: attribute SectionID must be positive")
	}

	// Aggregate accumulated errors
	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		allErrors := append(nilPointerErrors, validationErrors...)
		return fmt.Errorf("validation errors: %v", allErrors)
	}

	return nil
}

// Create handles the creation of a new product batch.
// @Summary Create a new product batch
// @Description Create a new product batch with the provided JSON data
// @Tags product_batches
// @Accept json
// @Produce json
// @Param productBatch body NewProductBatchesReqJSON true "Product Batch Create JSON"
// @Success 201 {object} ProductBatchesResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/product_batches [post]
func (ctl *ProductBatchesController) Create(w http.ResponseWriter, r *http.Request) {
	var prodBatchReqJSON NewProductBatchesReqJSON
	if err := json.NewDecoder(r.Body).Decode(&prodBatchReqJSON); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}

	if err := prodBatchReqJSON.validateCreate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusUnprocessableEntity, err.Error()))

		return
	}

	prodBatchDTO := models.ProductBatchesDTO{
		BatchNumber:        *prodBatchReqJSON.BatchNumber,
		InitialQuantity:    *prodBatchReqJSON.InitialQuantity,
		CurrentQuantity:    *prodBatchReqJSON.CurrentQuantity,
		CurrentTemperature: *prodBatchReqJSON.CurrentTemperature,
		MinimumTemperature: *prodBatchReqJSON.MinimumTemperature,
		DueDate:            *prodBatchReqJSON.DueDate,
		ManufacturingDate:  *prodBatchReqJSON.ManufacturingDate,
		ManufacturingHour:  *prodBatchReqJSON.ManufacturingHour,
		ProductID:          *prodBatchReqJSON.ProductID,
		SectionID:          *prodBatchReqJSON.SectionID,
	}

	newProdBatch, err := ctl.sv.Create(prodBatchDTO)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok &&
			(mysqlErr.Number == mysqlerr.CodeDuplicateEntry ||
				mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow) {
			response.Error(w, http.StatusConflict, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusConflict, err.Error()))

			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	data := ProductBatchFullJSON{
		ID:                 newProdBatch.ID,
		BatchNumber:        newProdBatch.BatchNumber,
		InitialQuantity:    newProdBatch.InitialQuantity,
		CurrentQuantity:    newProdBatch.CurrentQuantity,
		CurrentTemperature: newProdBatch.CurrentTemperature,
		MinimumTemperature: newProdBatch.MinimumTemperature,
		DueDate:            newProdBatch.DueDate,
		ManufacturingDate:  newProdBatch.ManufacturingDate,
		ManufacturingHour:  newProdBatch.ManufacturingHour,
		ProductID:          newProdBatch.ProductID,
		SectionID:          newProdBatch.SectionID,
	}

	res := ProductBatchesResJSON{
		Message: "Created",
		Data:    data,
	}

	response.JSON(w, http.StatusCreated, res)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusCreated, res))
}
