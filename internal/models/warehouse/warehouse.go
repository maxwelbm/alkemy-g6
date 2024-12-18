package models

type Warehouse struct {
	Id					int		`json:"id"`
	Address				string	`json:"address"`
	Telephone			string	`json:"telephone"`
	WarehouseCode		string	`json:"warehouse_code"`
	MinimumCapacity		int		`json:"minimun_capacity"`
	MinimumTemperature	float64	`json:"minimum_temperature"`
}

type WarehouseDTO struct {
	Address				string
	Telephone			string
	WarehouseCode		string
	MinimumCapacity		int
	MinimumTemperature	float64
}