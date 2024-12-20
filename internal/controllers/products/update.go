package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) Update(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	// Decode the JSON
	var prodJson UpdateProductAttributesJSON
	json.NewDecoder(r.Body).Decode(&prodJson)

	// Validate the JSON
	if err = prodJson.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Create a new ProductDTO with the values from the JSON
	prodDTO := models.ProductDTO{}

	if prodJson.ProductCode != nil {
		prodDTO.ProductCode = *prodJson.ProductCode
	}
	if prodJson.Description != nil {
		prodDTO.Description = *prodJson.Description
	}
	if prodJson.Height != nil {
		prodDTO.Height = *prodJson.Height
	}
	if prodJson.Length != nil {
		prodDTO.Length = *prodJson.Length
	}
	if prodJson.Width != nil {
		prodDTO.Width = *prodJson.Width
	}
	if prodJson.Weight != nil {
		prodDTO.Weight = *prodJson.Weight
	}
	if prodJson.ExpirationRate != nil {
		prodDTO.ExpirationRate = *prodJson.ExpirationRate
	}
	if prodJson.FreezingRate != nil {
		prodDTO.FreezingRate = *prodJson.FreezingRate
	}
	if prodJson.RecomFreezTemp != nil {
		prodDTO.RecomFreezTemp = *prodJson.RecomFreezTemp
	}
	if prodJson.ProductTypeID != nil {
		prodDTO.ProductTypeID = *prodJson.ProductTypeID
	}
	if prodJson.SellerID != nil {
		prodDTO.SellerID = *prodJson.SellerID
	}

	// Update the product
	updatedProd, err := p.sv.Update(id, prodDTO)
	if errors.Is(err, repository.ErrProductNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	if errors.Is(err, repository.ErrProductUniqueness) {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Return the updated product
	res := ProductResJSON{Message: "Updated", Data: updatedProd}
	response.JSON(w, http.StatusOK, res)
}
