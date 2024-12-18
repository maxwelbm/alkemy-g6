package controller

import (
    "github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

type WarehouseDataResJSON struct {
    Id                  int     `json:"id"`
    Address             string  `json:"address"`
    Telephone           string  `json:"telephone"`
    WarehouseCode       string  `json:"warehouse_code"`
    MinimumCapacity     int     `json:"minimum_capacity"`
    MinimumTemperature  float64 `json:"minimum_temperature"`
}

type WarehouseResJSON struct {
    Message             string      `json:"message,omitempty"`
    Data                any         `json:"data,omitempty"`
}

type WarehouseReqJSON struct {
    Address             *string  `json:"address"`
    Telephone           *string  `json:"telephone"`
    WarehouseCode       *string  `json:"warehouse_code"`
    MinimumCapacity     *int     `json:"minimum_capacity"`
    MinimumTemperature  *float64 `json:"minimum_temperature"`
}

func NewWarehouseDefault(service models.WarehouseService) *WarehouseDefault {
    return &WarehouseDefault{service: service}
}

type WarehouseDefault struct {
    service models.WarehouseService
}