package sections_controller

import (
	"encoding/json"
	"net/http"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsController) Create(w http.ResponseWriter, r *http.Request) {
	var secReqJson NewSectionReqJSON
	err := json.NewDecoder(r.Body).Decode(&secReqJson)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = secReqJson.validateCreate()
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	secDTO := models.SectionDTO{
		SectionNumber:      secReqJson.SectionNumber,
		CurrentTemperature: secReqJson.CurrentTemperature,
		MinimumTemperature: secReqJson.MinimumTemperature,
		CurrentCapacity:    secReqJson.CurrentCapacity,
		MinimumCapacity:    secReqJson.MinimumCapacity,
		MaximumCapacity:    secReqJson.MaximumCapacity,
		WarehouseID:        secReqJson.WarehouseID,
		ProductTypeID:      secReqJson.ProductTypeID,
	}

	newSection, err := c.SV.Create(secDTO)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	res := SectionResJSON{
		Message: "Created",
		Data:    newSection,
	}
	response.JSON(w, http.StatusCreated, res)
	return
}
