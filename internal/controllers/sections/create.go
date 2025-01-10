package sections_controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsController) Create(w http.ResponseWriter, r *http.Request) {
	var secReqJson NewSectionReqJSON
	if err := json.NewDecoder(r.Body).Decode(&secReqJson); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := secReqJson.validateCreate(); err != nil {
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

	newSection, err := c.sv.Create(secDTO)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := SectionFullJSON{
		ID:                 newSection.ID,
		SectionNumber:      newSection.SectionNumber,
		CurrentTemperature: newSection.CurrentTemperature,
		MinimumTemperature: newSection.MinimumTemperature,
		CurrentCapacity:    newSection.CurrentCapacity,
		MinimumCapacity:    newSection.MinimumCapacity,
		MaximumCapacity:    newSection.MaximumCapacity,
		WarehouseID:        newSection.WarehouseID,
		ProductTypeID:      newSection.ProductTypeID,
	}

	res := SectionResJSON{
		Message: "Created",
		Data:    data,
	}

	response.JSON(w, http.StatusCreated, res)
}
