package localities_controller

import (
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

type LocalityReportJSON struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}

// ReportSellers retrieves the report of sellers for a given locality ID.
// @Summary Get locality sellers report
// @Description Retrieve the number of sellers in a locality by ID
// @Tags localities
// @Produce json
// @Param id query int true "Locality ID"
// @Success 200 {object} LocalityResJSON "OK"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /localities/report_sellers [get]
func (ct *LocalityController) ReportSellers(w http.ResponseWriter, r *http.Request) {
	// Extract the "id" parameter from the URL query and convert it to an integer
	var id int
	var err error
	paramsId := r.URL.Query().Get("id")
	if paramsId != "" {
		id, err = strconv.Atoi(paramsId)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		if id < 1 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}

	// Call the service layer to get the locality sellers report by id
	locs, err := ct.sv.ReportSellers(id)
	if err != nil {
		// If an error occurs, return an internal server error
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Populate the response JSON with the locality report data
	var data []LocalityReportJSON
	for _, loc := range locs {
		locJson := LocalityReportJSON{
			ID:           loc.ID,
			LocalityName: loc.LocalityName,
			SellersCount: loc.SellersCount,
		}
		data = append(data, locJson)
	}

	// Create the response JSON and send it with status OK
	var res LocalityResJSON
	if len(data) == 1 {
		res = LocalityResJSON{Message: "OK", Data: data[0]}
	} else {
		res = LocalityResJSON{Data: data}
	}
	response.JSON(w, http.StatusOK, res)
}
