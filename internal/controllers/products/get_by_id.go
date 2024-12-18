package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	prod, err := p.sv.GetById(id)
	if err == repository.ErrProductNotFound {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data = ProductFullJSON{
		ID:             prod.ID,
		ProductCode:    prod.ProductCode,
		Description:    prod.Description,
		Height:         prod.Height,
		Length:         prod.Length,
		Width:          prod.Width,
		Weight:         prod.Weight,
		ExpirationRate: prod.ExpirationRate,
		FreezingRate:   prod.FreezingRate,
		RecomFreezTemp: prod.RecomFreezTemp,
		ProductTypeID:  prod.ProductTypeID,
		SellerID:       prod.SellerID,
	}

	res := ProductResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
