package controllers

import (
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

type WarehouseResJSON struct {
	Address				string	`json:"address"`
	Telephone			string	`json:"telephone"`
	WarehouseCode		string	`json:"warehouse_code"`
	MinimumCapacity		int		`json:"minimun_capacity"`
	MinimumTemperature	float64	`json:"minimum_temperature"`
}

func NewWarehouseDefault(service models.WarehouseService) *WarehouseDefault {
	return &WarehouseDefault{service: service}
}

type WarehouseDefault struct {
	service models.WarehouseService
}