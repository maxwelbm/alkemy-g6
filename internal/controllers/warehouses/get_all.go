package controller

import (
    "net/http"
    "github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) GetAll(w http.ResponseWriter, r *http.Request) {
    v, err := c.service.GetAll()
    if err != nil {
        response.Error(w, http.StatusInternalServerError, "Failed to retrieve warehouses")
        return
    }

    var data []WarehouseDataResJSON
    for _, value := range v {
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
