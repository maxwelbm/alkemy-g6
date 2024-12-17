package controllers

import (
	"net/http"
)

func (h *WarehouseDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.GetAllWarehouses()
		if err != nil {
			// response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
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
		// response.JSON(w, http.StatusOK, map[string]any{
		// 	"message": "success",
		// 	"data":    data,
		// })
	}
}