package sections

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (h *SectionsDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sec, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		data := make(map[int]SectionFullJSON)
		for key, value := range sec {
			data[key] = SectionFullJSON{
				ID:                 value.ID,
				SectionNumber:      value.SectionNumber,
				CurrentTemperature: value.CurrentTemperature,
				MinimumTemperature: value.MinimumTemperature,
				CurrentCapacity:    value.CurrentCapacity,
				MinimumCapacity:    value.MinimumCapacity,
				MaximumCapacity:    value.MaximumCapacity,
				WarehouseID:        value.WarehouseID,
				ProductTypeID:      value.ProductTypeID,
			}
		}
		response.JSON(w, http.StatusOK, data)
	}
}
