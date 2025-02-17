package warehousesctl

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetByID handles the HTTP request to retrieve a warehouse by its ID.
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
func (c *WarehouseDefault) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}
	// If the ID is less than 1, return a 400 Bad Request error
	if id < 1 {
		response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))

		return
	}

	warehouse, err := c.sv.GetByID(id)
	if err != nil {
		// If the section is not found, return a 404 Not Found error
		if errors.Is(err, models.ErrWareHouseNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

			return
		}

		// If the section is not found, return a 500 Internal Server Error error
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	data := WarehouseDataResJSON{
		ID:                 warehouse.ID,
		Address:            warehouse.Address,
		Telephone:          warehouse.Telephone,
		WarehouseCode:      warehouse.WarehouseCode,
		MinimumCapacity:    warehouse.MinimumCapacity,
		MinimumTemperature: warehouse.MinimumTemperature,
	}
	res := WarehouseResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusOK, res))
}
