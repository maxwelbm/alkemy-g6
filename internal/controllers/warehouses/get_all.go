package controller

import (
	"net/http"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	v, err := c.service.GetAllWarehouses()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve warehouses")
		return
	}

	var data []WarehouseResJSON
	for _, value := range v {
		new := WarehouseResJSON{
			Id:                 value.Id,
			Address:           	value.Address,
			Telephone:          value.Telephone,
			WarehouseCode:    	value.WarehouseCode,
			MinimumCapacity:    value.MinimumCapacity,
			MinimumTemperature: value.MinimumTemperature,
		}
		data = append(data, new)
	}
	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}