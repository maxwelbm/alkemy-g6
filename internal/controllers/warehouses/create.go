package warehouses_controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var warehouseJSON WarehouseReqJSON
	err := json.NewDecoder(r.Body).Decode(&warehouseJSON)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateNewWarehouse(warehouseJSON)
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

	resWarehouse, err := c.Service.Create(warehouse)
	if err != nil {
		if err.(*mysql.MySQLError).Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
		}
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
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
	response.JSON(w, http.StatusCreated, res)
}
