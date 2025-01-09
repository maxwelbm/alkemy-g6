package sections_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var secReqJson NewSectionReqJSON
	if err = json.NewDecoder(r.Body).Decode(&secReqJson); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = secReqJson.validateUpdate(); err != nil {
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

	updateSection, err := c.SV.Update(id, secDTO)

	if err != nil {
		// Handle no changes made
		if errors.Is(err, models.ErrorNoChangesMade) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		// Handle MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		// Handle other internal server errors
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := SectionResJSON{
		Message: "Success",
		Data:    updateSection,
	}
	response.JSON(w, http.StatusOK, res)
}
