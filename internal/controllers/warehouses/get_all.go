package controller

import (
	"net/http"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v, err := c.service.GetAllWarehouses()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := make(map[int]WarehouseResJSON)
		for key, value := range v {
			data[key] = WarehouseResJSON{
				Address:           	value.Address,
				Telephone:          value.Telephone,
				WarehouseCode:    	value.WarehouseCode,
				MinimumCapacity:    value.MinimumCapacity,
				MinimumTemperature: value.MinimumTemperature,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}