package localitiesctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type NewLocalityJSON struct {
	LocalityName *string `json:"locality_name,omitempty"`
	ProvinceName *string `json:"province_name,omitempty"`
	CountryName  *string `json:"country_name,omitempty"`
}

func (j *NewLocalityJSON) validate() (err error) {
	var locErrs []string

	if j.LocalityName == nil {
		locErrs = append(locErrs, "error: locality_name cannot be nil")
	}

	if j.LocalityName != nil && *j.LocalityName == "" {
		locErrs = append(locErrs, "error: locality_name cannot be empty")
	}

	if j.ProvinceName == nil {
		locErrs = append(locErrs, "error: province_name cannot be nil")
	}

	if j.ProvinceName != nil && *j.ProvinceName == "" {
		locErrs = append(locErrs, "error: province_name cannot be empty")
	}

	if j.CountryName == nil {
		locErrs = append(locErrs, "error: country_name cannot be nil")
	}

	if j.CountryName != nil && *j.CountryName == "" {
		locErrs = append(locErrs, "error: country_name cannot be empty")
	}

	if len(locErrs) > 0 {
		err = fmt.Errorf("validation errors: %v", locErrs)
	}

	return err
}

// Create handles the creation of a new locality.
// @Summary Create a new locality
// @Description Create a new locality with the provided JSON payload
// @Tags localities
// @Accept json
// @Produce json
// @Param locality body NewLocalityJSON true "New Locality JSON"
// @Success 201 {object} models.LocalityDTO "Created"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/localities [post]
func (ct *LocalitiesController) Create(w http.ResponseWriter, r *http.Request) {
	// parse json
	var locJSON NewLocalityJSON
	err := json.NewDecoder(r.Body).Decode(&locJSON)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

		return
	}

	if err = locJSON.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusUnprocessableEntity, err.Error()))

		return
	}
	// builds dto from json
	locDTO := &models.LocalityDTO{
		LocalityName: locJSON.LocalityName,
		ProvinceName: locJSON.ProvinceName,
		CountryName:  locJSON.CountryName,
	}

	// insert
	loc, err := ct.sv.Create(*locDTO)
	if err != nil {
		// handles conflict error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusConflict, err.Error()))

			return
		}
		// handles other errors
		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	// builds json
	var localityJSON = FullLocalitySON{
		ID:           loc.ID,
		LocalityName: loc.LocalityName,
		ProvinceName: loc.ProvinceName,
		CountryName:  loc.CountryName,
	}

	// response
	data := LocalityResJSON{
		Message: http.StatusText(http.StatusCreated),
		Data:    localityJSON,
	}
	response.JSON(w, http.StatusCreated, data)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusCreated, data))
}
