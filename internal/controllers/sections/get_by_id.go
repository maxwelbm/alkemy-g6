package sectionsctl

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetByID handles the HTTP request to retrieve a section by its ID.
// @Summary Get a section by ID
// @Description Get a section by ID
// @Tags sections
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Success 200 {object} SectionFullJSON
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Router /api/v1/sections/{id} [get]
func (ctl *SectionsController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((chi.URLParam(r, "id")))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	sec, err := ctl.sv.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	data := SectionFullJSON{
		ID:                 sec.ID,
		SectionNumber:      sec.SectionNumber,
		CurrentTemperature: sec.CurrentTemperature,
		MinimumTemperature: sec.MinimumTemperature,
		CurrentCapacity:    sec.CurrentCapacity,
		MinimumCapacity:    sec.MinimumCapacity,
		MaximumCapacity:    sec.MaximumCapacity,
		WarehouseID:        sec.WarehouseID,
		ProductTypeID:      sec.ProductTypeID,
	}

	res := SectionResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}
