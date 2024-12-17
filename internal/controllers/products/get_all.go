package products_controller

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prods, err := p.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		data := make(map[int]ProductFullJSON)
		for k, v := range prods {
			data[k] = ProductFullJSON{
				ID:             v.ID,
				ProductCode:    v.ProductCode,
				Description:    v.Description,
				Height:         v.Height,
				Length:         v.Length,
				Width:          v.Width,
				Weight:         v.Weight,
				ExpirationRate: v.ExpirationRate,
				FreezingRate:   v.FreezingRate,
				RecomFreezTemp: v.RecomFreezTemp,
				ProductTypeID:  v.ProductTypeID,
				SellerID:       v.SellerID,
			}
		}

		response.JSON(w, http.StatusOK, data)
	}
}
