package sectionsctl

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/logger"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// ReportProducts handles the HTTP request to retrieve report products for a specific section.
// @Summary Retrieve report products for a section
// @Description Retrieves report products for a section based on the provided section ID.
// @Tags sections
// @Produce json
// @Param id query int true "Section ID"
// @Success 200 {array} ReportProductFullJSON "Successfully retrieved report products"
// @Failure 400 {object} response.ErrorResponse "Invalid section ID or bad request"
// @Failure 404 {object} response.ErrorResponse "Section not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/sections/reportProducts [get]
func (ctl *SectionsController) ReportProducts(w http.ResponseWriter, r *http.Request) {
	var sectionID int

	var err error

	id := r.URL.Query().Get("id")
	if id != "" {
		sectionID, err = strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, err.Error()))

			return
		}

		if sectionID < 1 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))

			return
		}
	}

	var reportProducts []models.ProductReport
	reportProducts, err = ctl.sv.ReportProducts(sectionID)

	if err != nil {
		if errors.Is(err, models.ErrSectionNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusNotFound, err.Error()))

			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		logger.Writer.Error(fmt.Sprintf("HTTP Status Code: %d - %s", http.StatusInternalServerError, err.Error()))

		return
	}

	data := make([]ReportProductFullJSON, 0, len(reportProducts))
	for _, r := range reportProducts {
		data = append(data,
			ReportProductFullJSON{
				SectionID:     r.SectionID,
				SectionNumber: r.SectionNumber,
				ProductsCount: r.ProductsCount,
			})
	}

	res := ProductReportResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
	logger.Writer.Info(fmt.Sprintf("HTTP Status Code: %d - %#v", http.StatusOK, res))
}
