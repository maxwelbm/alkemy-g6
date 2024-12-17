package products_controller

import (
	"encoding/json"
	"net/http"
)

func (p *ProductsDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prods, err := p.sv.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(`{"message": "badbad :("}`)
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}
