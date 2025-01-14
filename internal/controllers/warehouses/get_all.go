package warehousesctl

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetAll handles the HTTP GET request to retrieve all warehouses.
//
// @Summary Retrieve all warehouses
// @Description This endpoint retrieves all warehouses from the database using the service layer.
// @Tags warehouses
// @Produce json
// @Success 200 {object} WarehouseResJSON "OK - The warehouses were successfully retrieved"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error - An unexpected error occurred during the retrieval process"
// @Router /api/v1/warehouses [get]
func (c *WarehouseDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	warehouses, err := c.sv.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := make([]WarehouseDataResJSON, 0, len(warehouses))

	for _, value := range warehouses {
		warehouse := WarehouseDataResJSON{
			ID:                 value.ID,
			Address:            value.Address,
			Telephone:          value.Telephone,
			WarehouseCode:      value.WarehouseCode,
			MinimumCapacity:    value.MinimumCapacity,
			MinimumTemperature: value.MinimumTemperature,
		}
		data = append(data, warehouse)
	}

	res := WarehouseResJSON{Data: data}

	response.JSON(w, http.StatusOK, res)
}
