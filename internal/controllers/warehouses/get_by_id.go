package controller

import (
	"errors"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) GetById(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        response.Error(w, http.StatusBadRequest, err.Error())
        return
    }
    warehouse, err := c.service.GetById(id)
    if errors.Is(err, models.ErrWarehouseRepositoryNotFound) {
        response.Error(w, http.StatusNotFound, err.Error())
        return
    }
    if err != nil {
        response.JSON(w, http.StatusInternalServerError, nil)
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
