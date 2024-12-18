package controller

import (
	"encoding/json"
	"net/http"

	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) Create(w http.ResponseWriter, r *http.Request) {
	var prodJson NewProductAttributesJSON
	json.NewDecoder(r.Body).Decode(&prodJson)

	err := prodJson.validate()
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	prodDTO := models.ProductDTO{
		ProductCode:    *prodJson.ProductCode,
		Description:    *prodJson.Description,
		Height:         *prodJson.Height,
		Length:         *prodJson.Length,
		Width:          *prodJson.Width,
		Weight:         *prodJson.Weight,
		ExpirationRate: *prodJson.ExpirationRate,
		FreezingRate:   *prodJson.FreezingRate,
		RecomFreezTemp: *prodJson.RecomFreezTemp,
		ProductTypeID:  *prodJson.ProductTypeID,
	}
	if prodJson.SellerID != nil {
		prodDTO.SellerID = *prodJson.SellerID
	}

	newProd, err := p.sv.Create(prodDTO)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	res := ProductResJSON{Message: "Created", Data: newProd}
	response.JSON(w, http.StatusCreated, res)
}
