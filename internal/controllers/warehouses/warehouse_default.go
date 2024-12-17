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

func NewWarehouseDefault(sv models.WarehouseService) *WarehouseDefault {
	return &WarehouseDefault{sv: sv}
}

type WarehouseDefault struct {
	sv models.WarehouseService
}