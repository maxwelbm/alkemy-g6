package sections

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsDefault) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}
	var secReqJson NewSectionReqJSON
	err = json.NewDecoder(r.Body).Decode(&secReqJson)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, nil)
		return
	}
	err = secReqJson.validateUpdate()
	// err = validateUpdateWarehouse(secReqJson)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
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

	updateSection, err := c.sv.Update(id, secDTO)
	if errors.Is(err, repository.ErrSectionDuplicatedCode) {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}
	if errors.Is(err, repository.ErrSectionNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	res := SectionResJSON{
		Message: "Success",
		Data:    updateSection,
	}
	response.JSON(w, http.StatusOK, res)
	return
}
