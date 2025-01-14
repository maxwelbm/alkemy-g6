package productsctl

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
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
func (p *ProductsDefault) Update(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Decode the JSON
	var prodJSON UpdateProductAttributesJSON
	json.NewDecoder(r.Body).Decode(&prodJSON)

	// Validate the JSON
	if err = prodJSON.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Create a new ProductDTO with the values from the JSON
	prodDTO := models.ProductDTO{}

	if prodJSON.ProductCode != nil {
		prodDTO.ProductCode = *prodJSON.ProductCode
	}
	if prodJSON.Description != nil {
		prodDTO.Description = *prodJSON.Description
	}
	if prodJSON.Height != nil {
		prodDTO.Height = *prodJSON.Height
	}
	if prodJSON.Length != nil {
		prodDTO.Length = *prodJSON.Length
	}
	if prodJSON.Width != nil {
		prodDTO.Width = *prodJSON.Width
	}
	if prodJSON.NetWeight != nil {
		prodDTO.NetWeight = *prodJSON.NetWeight
	}
	if prodJSON.ExpirationRate != nil {
		prodDTO.ExpirationRate = *prodJSON.ExpirationRate
	}
	if prodJSON.FreezingRate != nil {
		prodDTO.FreezingRate = *prodJSON.FreezingRate
	}
	if prodJSON.RecomFreezTemp != nil {
		prodDTO.RecomFreezTemp = *prodJSON.RecomFreezTemp
	}
	if prodJSON.ProductTypeID != nil {
		prodDTO.ProductTypeID = *prodJSON.ProductTypeID
	}
	if prodJSON.SellerID != nil {
		prodDTO.SellerID = *prodJSON.SellerID
	}

	// Update the product
	updatedProd, err := p.SV.Update(id, prodDTO)

	if errors.Is(err, models.ErrProductNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case mysqlerr.CodeDuplicateEntry:
				response.Error(w, http.StatusConflict, err.Error())
				return
			case mysqlerr.CodeCannotAddOrUpdateChildRow:
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return the updated product
	res := ProductResJSON{Message: "Updated", Data: updatedProd}
	response.JSON(w, http.StatusOK, res)
}
