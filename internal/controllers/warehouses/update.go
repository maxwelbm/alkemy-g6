package controller

import (
	"encoding/json"
	"errors"
	"net/http"
    "strconv"
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
    "github.com/go-chi/chi/v5"
)

func (c *WarehouseDefault) Update(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid ID format")
        return
    }
    var warehouseJSON WarehouseReqJSON
    err = json.NewDecoder(r.Body).Decode(&warehouseJSON)
    if err != nil {
        response.JSON(w, http.StatusBadRequest, nil)
        return
    }
    err = validateUpdateWarehouse(warehouseJSON)
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
    wh, err := c.service.Update(id, warehouse)
    if errors.Is(err, repository.ErrWarehouseRepositoryDuplicatedCode) {
        response.Error(w, http.StatusConflict, err.Error())
        return
    }
    if errors.Is(err, repository.ErrWarehouseRepositoryNotFound) {
        response.Error(w, http.StatusNotFound, err.Error())
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
    response.JSON(w, http.StatusOK, res)
}
