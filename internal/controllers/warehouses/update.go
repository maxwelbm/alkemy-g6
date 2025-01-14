package warehousesctl

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Update handles the HTTP PUT request to update a warehouse by its ID.
// It extracts the warehouse ID from the URL parameters, validates it,
// and fetches the warehouse details from the service layer. If the ID
// is invalid or the warehouse is not found, it returns an appropriate
// error response.
//
// @Summary Update a warehouse
// @Description Update a warehouse by its ID
// @Tags warehouses
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Param data body WarehouseReqJSON true "Warehouse JSON"
// @Success 200 {object} WarehouseResJSON "OK - The warehouse was successfully updated"
// @Failure 400 {object} response.ErrorResponse "Bad request - invalid ID format or incomplete data"
// @Failure 404 {object} response.ErrorResponse "Not found - the warehouse does not exist"
// @Failure 409 {object} response.ErrorResponse "Conflict - duplicate entry"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/warehouses/{id} [patch]
func (c *WarehouseDefault) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var warehouseJSON WarehouseReqJSON
	err = json.NewDecoder(r.Body).Decode(&warehouseJSON)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateUpdateWarehouse(warehouseJSON)
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

	resWarehouse, err := c.sv.Update(id, warehouse)
	if err != nil {
		if errors.Is(err, models.ErrWareHouseCodeExist) {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}

		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

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
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}
