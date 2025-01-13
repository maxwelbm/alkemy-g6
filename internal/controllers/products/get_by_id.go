package products_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetById retrieves a product by its ID.
//
// @Summary Get product by ID
// @Description Get a product by its unique ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} ProductResJSON "Product found"
// @Failure 400 {object} ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /products/{id} [get]
func (p *ProductsDefault) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	prod, err := p.SV.GetById(id)
	if errors.Is(err, models.ErrProductNotFound) {
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
		NetWeight:         prod.NetWeight,
		ExpirationRate: prod.ExpirationRate,
		FreezingRate:   prod.FreezingRate,
		RecomFreezTemp: prod.RecomFreezTemp,
		ProductTypeID:  prod.ProductTypeID,
		SellerID:       prod.SellerID,
	}

	res := ProductResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
