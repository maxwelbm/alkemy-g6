package sections_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsController) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((chi.URLParam(r, "id")))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	sec, err := c.SV.GetById(id)
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
