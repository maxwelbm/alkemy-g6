package controller

import (
    "github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

type WarehouseDataResJSON struct {
    Id                  int     `json:"id"`
    Address             string  `json:"address"`
    Telephone           string  `json:"telephone"`
    WarehouseCode       string  `json:"warehouse_code"`
    MinimumCapacity     int     `json:"minimun_capacity"`
    MinimumTemperature  float64 `json:"minimum_temperature"`
}

type WarehousesResJSON struct {
    Message             string                  `json:"message"`
    Data                []WarehouseDataResJSON  `json:"data"`
}

type WarehouseResJSON struct {
    Message             string                  `json:"message"`
    Data                WarehouseDataResJSON    `json:"data"`
}

func NewWarehouseDefault(service models.WarehouseService) *WarehouseDefault {
    return &WarehouseDefault{service: service}
}

type WarehouseDefault struct {
    service models.WarehouseService
}