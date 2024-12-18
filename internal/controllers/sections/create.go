package sections

import (
	"encoding/json"
	"errors"
	"net/http"

	models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsDefault) Create(w http.ResponseWriter, r *http.Request) {
	var secReqJson NewSectionReqJSON
	json.NewDecoder(r.Body).Decode(&secReqJson)

	err := secReqJson.validate()
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	secDTO := models.SectionDTO{
		SectionNumber:      *secReqJson.SectionNumber,
		CurrentTemperature: *secReqJson.CurrentTemperature,
		MinimumTemperature: *secReqJson.MinimumTemperature,
		CurrentCapacity:    *secReqJson.CurrentCapacity,
		MinimumCapacity:    *secReqJson.MinimumCapacity,
		MaximumCapacity:    *secReqJson.MaximumCapacity,
		WarehouseID:        *secReqJson.WarehouseID,
		ProductTypeID:      *secReqJson.ProductTypeID,
	}

	newSection, err := c.sv.Create(secDTO)
	if errors.Is(err, repository.ErrSectionDuplicatedCode) {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}
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
