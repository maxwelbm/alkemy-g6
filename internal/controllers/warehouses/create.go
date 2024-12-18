package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
	"github.com/maxwelbm/alkemy-g6/internal/repository/warehouse"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func validateNewWarehouse(warehouse WarehouseReqJSON) error {
    if warehouse.Address == nil {
        return errors.New("The address field cannot be empty")
    }
    if warehouse.Telephone == nil {
        return errors.New("The telephone field cannot be empty")
    }
    if warehouse.WarehouseCode == nil {
        return errors.New("The warehouse_code field cannot be empty")
    }
    if warehouse.MinimumCapacity == nil {
        return errors.New("The minimum_capacity field cannot be empty")
    }
    if *warehouse.MinimumCapacity <= 0 {
        return errors.New("The minimum_capacity field must be greater than zero")
    }
    if warehouse.MinimumTemperature == nil {
        return errors.New("The minimum_temperature field cannot be empty")
    }

    return nil
}

func (c *WarehouseDefault) Create(w http.ResponseWriter, r *http.Request) {
    var warehouseJSON WarehouseReqJSON
    err := json.NewDecoder(r.Body).Decode(&warehouseJSON)
    if err != nil {
        response.JSON(w, http.StatusBadRequest, nil)
        return
    }
    err = validateNewWarehouse(warehouseJSON)
    if err != nil {
        response.JSON(w, http.StatusUnprocessableEntity, err.Error())
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
