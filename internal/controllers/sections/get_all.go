package sectionsctl

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// @Summary Get all sections
// @Description Get all sections
// @Tags sections
// @Accept json
// @Produce json
// @Success 200 {array} SectionFullJSON
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/sections [get]
func (c *SectionsController) GetAll(w http.ResponseWriter, r *http.Request) {
	sec, err := c.sv.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data []SectionFullJSON
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
