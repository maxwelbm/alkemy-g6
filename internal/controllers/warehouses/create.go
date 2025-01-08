package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) Create(w http.ResponseWriter, r *http.Request) {
    var warehouseJSON WarehouseReqJSON
    err := json.NewDecoder(r.Body).Decode(&warehouseJSON)
    if err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid request body format")
        return
    }
    err = validateNewWarehouse(warehouseJSON)
    if err != nil {
        response.Error(w, http.StatusUnprocessableEntity, err.Error())
        return
    }
    warehouse := models.WarehouseDTO{
        Address: warehouseJSON.Address,
        Telephone: warehouseJSON.Telephone,
        WarehouseCode: warehouseJSON.WarehouseCode,
        MinimumCapacity: warehouseJSON.MinimumCapacity,
        MinimumTemperature: warehouseJSON.MinimumTemperature,
    }
    wh, err := c.service.Create(warehouse)
    if errors.Is(err, repository.ErrWarehouseRepositoryDuplicatedCode) {
        response.Error(w, http.StatusConflict, err.Error())
        return
    }
    if err != nil {
        response.JSON(w, http.StatusUnprocessableEntity, err.Error())
        return
    }

    data := WarehouseDataResJSON{
        Id:                 wh.Id,
        Address:            wh.Address,
        Telephone:          wh.Telephone,
        WarehouseCode:      wh.WarehouseCode,
        MinimumCapacity:    wh.MinimumCapacity,
        MinimumTemperature: wh.MinimumTemperature,
    }
    res := WarehouseResJSON{
        Message: "Success",
        Data:    data,
    }
    response.JSON(w, http.StatusCreated, res)
}
