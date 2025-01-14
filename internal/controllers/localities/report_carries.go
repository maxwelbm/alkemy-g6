package localitiesctl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type CarryReportJSON struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int    `json:"carries_count"`
}

// ReportCarries retrieves the report of carries for a given locality ID.
// @Summary Get locality carries report
// @Description Retrieve the number of carries in a locality by ID
// @Tags localities
// @Produce json
// @Param id query int true "Carry ID"
// @Success 200 {object} CarryReportJSON "OK"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 404 {object} response.ErrorResponse "Not Found"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/localities/reportCarries [get]
func (ct *LocalitiesController) ReportCarries(w http.ResponseWriter, r *http.Request) {
	// Extract the "id" parameter from the URL query and convert it to an integer
	var id int

	var err error

	paramsID := r.URL.Query().Get("id")
	if paramsID != "" {
		id, err = strconv.Atoi(paramsID)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if id < 1 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}

	// Call the service layer to get the locality carries report by id
	locs, err := ct.sv.ReportCarries(id)
	if err != nil {
		// If an id does not exist, return status not found
		if errors.Is(err, models.ErrCarryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		// If an error occurs, return an internal server error
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	// Populate the response JSON with the locality report data
	data := make([]CarryReportJSON, len(locs))

	for _, loc := range locs {
		locJSON := CarryReportJSON{
			ID:           loc.ID,
			LocalityName: loc.LocalityName,
			CarriesCount: loc.CarriesCount,
		}
		data = append(data, locJSON)
	}

	// Create the response JSON and send it with status OK
	var res LocalityResJSON
	res = LocalityResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
