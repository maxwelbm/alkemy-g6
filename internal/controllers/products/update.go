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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var prodJson UpdateProductAttributesJSON
	json.NewDecoder(r.Body).Decode(&prodJson)

	prodDTO := models.ProductDTO{
		ProductCode:    prodJson.ProductCode,
		Description:    prodJson.Description,
		Height:         prodJson.Height,
		Length:         prodJson.Length,
		Width:          prodJson.Width,
		Weight:         prodJson.Weight,
		ExpirationRate: prodJson.ExpirationRate,
		FreezingRate:   prodJson.FreezingRate,
		RecomFreezTemp: prodJson.RecomFreezTemp,
		ProductTypeID:  prodJson.ProductTypeID,
		SellerID:       prodJson.SellerID,
	}

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

	res := ProductResJSON{Message: "Updated", Data: updatedProd}
	response.JSON(w, http.StatusOK, res)
}
