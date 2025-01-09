package warehouses_controller

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	warehouses, err := c.Service.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data []WarehouseDataResJSON
	for _, value := range warehouses {
		new := WarehouseDataResJSON{
			Id:                 value.Id,
			Address:            value.Address,
			Telephone:          value.Telephone,
			WarehouseCode:      value.WarehouseCode,
			MinimumCapacity:    value.MinimumCapacity,
			MinimumTemperature: value.MinimumTemperature,
		}
		data = append(data, new)
	}

	res := WarehouseResJSON{
		Message: "Success",
		Data:    data,
	}

	response.JSON(w, http.StatusOK, res)
}
