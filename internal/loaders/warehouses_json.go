package loaders

import (
	"encoding/json"
	"os"   
    "github.com/maxwelbm/alkemy-g6/internal/models/warehouse"
)

func NewWarehouseJSONFile(path string) *WarehouseJSONFile {
    return &WarehouseJSONFile{
        path: path,
    }
}

type WarehouseJSONFile struct {
    path string
}

type WarehouseJSON struct {
    Id                	int     `json:"id"`
    WarehouseCode		string	`json:"warehouse_code"`
    Address				string  `json:"address"`
    Telephone			string  `json:"telephone"`
    MinimumCapacity		int     `json:"minimum_capacity"`
    MinimumTemperature	float64 `json:"minimum_temperature"`
}

func (l *WarehouseJSONFile) Load() (warehouses map[int]models.Warehouse, err error) {
    // open file
    file, err := os.Open(l.path)
    if err != nil {
        return
    }
    defer file.Close()

    // decode file
    var warehousesJSON []WarehouseJSON
    err = json.NewDecoder(file).Decode(&warehousesJSON)
    if err != nil {
        return
    }

    // serialize warehouses
    warehouses = make(map[int]models.Warehouse)
    for _, w := range warehousesJSON {
        warehouses[w.Id] = models.Warehouse{
            Id:                 w.Id,
            WarehouseCode:      w.WarehouseCode,
            Address: 			w.Address,
			Telephone: 			w.Telephone,
            MinimumCapacity:    w.MinimumCapacity,
            MinimumTemperature: w.MinimumTemperature,
        }
    }

    return
}