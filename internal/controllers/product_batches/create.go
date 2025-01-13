package product_batches_controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *ProductBatchesController) Create(w http.ResponseWriter, r *http.Request) {
	var prodBatchReqJson NewProductBatchesReqJSON
	if err := json.NewDecoder(r.Body).Decode(&prodBatchReqJson); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := prodBatchReqJson.validateCreate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	prodBatchDTO := models.ProductBatchesDTO{
		BatchNumber:        *prodBatchReqJson.BatchNumber,
		InitialQuantity:    *prodBatchReqJson.InitialQuantity,
		CurrentQuantity:    *prodBatchReqJson.InitialQuantity,
		CurrentTemperature: *prodBatchReqJson.CurrentTemperature,
		MinimumTemperature: *prodBatchReqJson.MinimumTemperature,
		DueDate:            *prodBatchReqJson.DueDate,
		ManufacturingDate:  *prodBatchReqJson.ManufacturingDate,
		ManufacturingHour:  *prodBatchReqJson.ManufacturingHour,
		ProductID:          *prodBatchReqJson.ProductID,
		SectionID:          *prodBatchReqJson.SectionID,
	}

	newProdBatch, err := c.sv.Create(prodBatchDTO)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
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
		CurrentQuantity:    newProdBatch.InitialQuantity,
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
