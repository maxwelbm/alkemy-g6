package sectionsctl

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetAll handles the HTTP request to retrieve all sections.
// @Summary Get all sections
// @Description Get all sections
// @Tags sections
// @Accept json
// @Produce json
// @Success 200 {array} SectionFullJSON
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/sections [get]
func (ctl *SectionsController) GetAll(w http.ResponseWriter, r *http.Request) {
	sec, err := ctl.sv.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := make([]SectionFullJSON, 0, len(sec))
	for _, value := range sec {
		data = append(data, SectionFullJSON{
			ID:                 value.ID,
			SectionNumber:      value.SectionNumber,
			CurrentTemperature: value.CurrentTemperature,
			MinimumTemperature: value.MinimumTemperature,
			CurrentCapacity:    value.CurrentCapacity,
			MinimumCapacity:    value.MinimumCapacity,
			MaximumCapacity:    value.MaximumCapacity,
			WarehouseID:        value.WarehouseID,
			ProductTypeID:      value.ProductTypeID,
		})
	}

	res := SectionResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)
}
