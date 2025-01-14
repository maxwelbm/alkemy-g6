package productbatchesctl

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Create handles the creation of a new product batch.
// @Summary Create a new product batch
// @Description Create a new product batch with the provided JSON data
// @Tags product_batches
// @Accept json
// @Produce json
// @Param productBatch body NewProductBatchesReqJSON true "Product Batch Create JSON"
// @Success 201 {object} ProductBatchesResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Error ao decodificar JSON"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/product_batches [post]
func (c *ProductBatchesController) Create(w http.ResponseWriter, r *http.Request) {
	var prodBatchReqJSON NewProductBatchesReqJSON
	if err := json.NewDecoder(r.Body).Decode(&prodBatchReqJSON); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := prodBatchReqJSON.validateCreate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
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

	newProdBatch, err := c.sv.Create(prodBatchDTO)
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
}
