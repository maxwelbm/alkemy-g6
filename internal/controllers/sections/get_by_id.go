package sections

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsDefault) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((chi.URLParam(r, "id")))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid Id format ")
		return
	}

	sec, err := c.sv.GetById(id)

	if err == repository.ErrSectionNotFound {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
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
