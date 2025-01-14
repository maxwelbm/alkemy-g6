package warehousesctl

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Create handles the creation of a new warehouse.
// @Summary Create a new warehouse
// @Description Creates a new warehouse with the provided information.
// @Tags warehouses
// @Accept json
// @Produce json
// @Param data body WarehouseReqJSON true "Warehouse JSON"
// @Success 201 {object} WarehouseResJSON "Created"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 409 {object} response.ErrorResponse "Conflict - duplicate entry"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/warehouses [post]
func (c *WarehouseDefault) Create(w http.ResponseWriter, r *http.Request) {
	var warehouseJSON WarehouseReqJSON

	err := json.NewDecoder(r.Body).Decode(&warehouseJSON)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateNewWarehouse(warehouseJSON)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	warehouse := models.WarehouseDTO{
		Address:            warehouseJSON.Address,
		Telephone:          warehouseJSON.Telephone,
		WarehouseCode:      warehouseJSON.WarehouseCode,
		MinimumCapacity:    warehouseJSON.MinimumCapacity,
		MinimumTemperature: warehouseJSON.MinimumTemperature,
	}

	resWarehouse, err := c.sv.Create(warehouse)
	if err != nil {
		if err.(*mysql.MySQLError).Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
		}

		response.JSON(w, http.StatusUnprocessableEntity, err.Error())

		return
	}

	data := WarehouseDataResJSON{
		ID:                 resWarehouse.ID,
		Address:            resWarehouse.Address,
		Telephone:          resWarehouse.Telephone,
		WarehouseCode:      resWarehouse.WarehouseCode,
		MinimumCapacity:    resWarehouse.MinimumCapacity,
		MinimumTemperature: resWarehouse.MinimumTemperature,
	}

	res := WarehouseResJSON{
		Message: http.StatusText(http.StatusCreated),
		Data:    data,
	}
	response.JSON(w, http.StatusCreated, res)
}
