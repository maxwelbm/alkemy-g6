package controller

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	prods, err := p.sv.GetAll()
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var data []ProductFullJSON
	for _, p := range prods {
		data = append(data,
			ProductFullJSON{
				ID:             p.ID,
				ProductCode:    p.ProductCode,
				Description:    p.Description,
				Height:         p.Height,
				Length:         p.Length,
				Width:          p.Width,
				Weight:         p.Weight,
				ExpirationRate: p.ExpirationRate,
				FreezingRate:   p.FreezingRate,
				RecomFreezTemp: p.RecomFreezTemp,
				ProductTypeID:  p.ProductTypeID,
				SellerID:       p.SellerID,
			})
	}

	response.JSON(w, http.StatusOK, data)
}
