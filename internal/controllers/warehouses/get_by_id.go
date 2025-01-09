package warehouses_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	warehouse, err := c.Service.GetById(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	data := WarehouseDataResJSON{
		Id:                 warehouse.Id,
		Address:            warehouse.Address,
		Telephone:          warehouse.Telephone,
		WarehouseCode:      warehouse.WarehouseCode,
		MinimumCapacity:    warehouse.MinimumCapacity,
		MinimumTemperature: warehouse.MinimumTemperature,
	}
	res := WarehouseResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}
