package productsctl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Update handles the HTTP request to update an existing product by its ID.
//
// @Summary Update a product
// @Description Update an existing product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body UpdateProductAttributesJSON true "Product attributes to update"
// @Success 200 {object} ProductResJSON "Updated product"
// @Failure 400 {object} response.ErrorResponse "Invalid request parameters"
// @Failure 404 {object} response.ErrorResponse "Product not found"
// @Failure 409 {object} response.ErrorResponse "Duplicate entry"
// @Failure 422 {object} response.ErrorResponse "Unprocessable entity"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id} [patch]
//
//nolint:all
func (p *ProductsDefault) Update(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}

	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))

		return
	}

	// Decode the JSON
	var prodJSON UpdateProductAttributesJSON
	if err := json.NewDecoder(r.Body).Decode(&prodJSON); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}
	// Validate the JSON
	if err = prodJSON.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusUnprocessableEntity, err.Error()))

		return
	}

	// Create a new ProductDTO with the values from the JSON
	prodDTO := models.ProductDTO{}

	if prodJSON.ProductCode != nil {
		prodDTO.ProductCode = prodJSON.ProductCode
	}

	if prodJSON.Description != nil {
		prodDTO.Description = prodJSON.Description
	}

	if prodJSON.Height != nil {
		prodDTO.Height = prodJSON.Height
	}

	if prodJSON.Length != nil {
		prodDTO.Length = prodJSON.Length
	}

	if prodJSON.Width != nil {
		prodDTO.Width = prodJSON.Width
	}

	if prodJSON.NetWeight != nil {
		prodDTO.NetWeight = prodJSON.NetWeight
	}

	if prodJSON.ExpirationRate != nil {
		prodDTO.ExpirationRate = prodJSON.ExpirationRate
	}

	if prodJSON.FreezingRate != nil {
		prodDTO.FreezingRate = prodJSON.FreezingRate
	}

	if prodJSON.RecomFreezTemp != nil {
		prodDTO.RecomFreezTemp = prodJSON.RecomFreezTemp
	}

	if prodJSON.ProductTypeID != nil {
		prodDTO.ProductTypeID = prodJSON.ProductTypeID
	}

	if prodJSON.SellerID != nil {
		prodDTO.SellerID = prodJSON.SellerID
	}
	// Update the product
	updatedProd, err := p.SV.Update(id, prodDTO)

	if err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

			return
		}

		if mysqlErr, ok := err.(*mysql.MySQLError); ok &&
			(mysqlErr.Number == mysqlerr.CodeDuplicateEntry ||
				mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow) {
			response.Error(w, http.StatusConflict, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusConflict, err.Error()))

			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	data := ProductFullJSON{
		ID:             updatedProd.ID,
		ProductCode:    updatedProd.ProductCode,
		Description:    updatedProd.Description,
		Height:         updatedProd.Height,
		Length:         updatedProd.Length,
		Width:          updatedProd.Width,
		NetWeight:      updatedProd.NetWeight,
		ExpirationRate: updatedProd.ExpirationRate,
		FreezingRate:   updatedProd.FreezingRate,
		RecomFreezTemp: updatedProd.RecomFreezTemp,
		ProductTypeID:  updatedProd.ProductTypeID,
		SellerID:       updatedProd.SellerID,
	}

	// Return the updated product
	res := ProductResJSON{
		Message: "Updated",
		Data:    data,
	}

	response.JSON(w, http.StatusOK, res)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusOK, res))
}
