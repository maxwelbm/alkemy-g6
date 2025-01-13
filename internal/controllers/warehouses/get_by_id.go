package warehouses_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetById handles the HTTP request to retrieve a warehouse by its ID.
// It extracts the warehouse ID from the URL parameters, validates it,
// and fetches the warehouse details from the service layer. If the ID
// is invalid or the warehouse is not found, it returns an appropriate
// error response.
//
// @Summary Get warehouse by ID
// @Description Retrieve warehouse details by its ID
// @Tags warehouses
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} WarehouseResJSON "Success"
// @Failure 400 {object} response.ErrorResponse "Invalid ID format"
// @Failure 404 {object} response.ErrorResponse "Warehouse not found"
// @Router /api/v1/warehouses/{id} [get]
func (c *WarehouseDefault) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	warehouse, err := c.sv.GetById(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
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
