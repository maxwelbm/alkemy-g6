package warehouses_controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	var warehouseJSON WarehouseReqJSON
	err = json.NewDecoder(r.Body).Decode(&warehouseJSON)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateUpdateWarehouse(warehouseJSON)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	warehouse := models.WarehouseDTO{
		Address:            warehouseJSON.Address,
		Telephone:          warehouseJSON.Telephone,
		WarehouseCode:      warehouseJSON.WarehouseCode,
		MinimumCapacity:    warehouseJSON.MinimumCapacity,
		MinimumTemperature: warehouseJSON.MinimumTemperature,
	}

	resWarehouse, err := c.Service.Update(id, warehouse)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
	}

	data := WarehouseDataResJSON{
		Id:                 resWarehouse.Id,
		Address:            resWarehouse.Address,
		Telephone:          resWarehouse.Telephone,
		WarehouseCode:      resWarehouse.WarehouseCode,
		MinimumCapacity:    resWarehouse.MinimumCapacity,
		MinimumTemperature: resWarehouse.MinimumTemperature,
	}
	res := WarehouseResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}
