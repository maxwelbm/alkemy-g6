package sections_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetReportProducts handles the HTTP request to retrieve report products for a specific section.
// @Summary Retrieve report products for a section
// @Description Retrieves report products for a section based on the provided section ID.
// @Tags sections
// @Produce json
// @Param id query int true "Section ID"
// @Success 200 {array} ReportProductFullJSON "Successfully retrieved report products"
// @Failure 400 {object} response.ErrorResponse "Invalid section ID or bad request"
// @Failure 404 {object} response.ErrorResponse "Section not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/sections/report-products [get]
func (p *SectionsController) GetReportProducts(w http.ResponseWriter, r *http.Request) {
	var sectionId int
	var err error

	id := r.URL.Query().Get("id")
	if id != "" {
		sectionId, err = strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		if sectionId < 1 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}

	var reportProducts []models.ProductReport
	reportProducts, err = p.sv.GetReportProducts(sectionId)

	if err != nil {
		if errors.Is(err, models.ErrSectionNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data []ReportProductFullJSON
	for _, r := range reportProducts {
		data = append(data,
			ReportProductFullJSON{
				SectionId:     r.SectionId,
				SectionNumber: r.SectionNumber,
				ProductsCount: r.ProductsCount,
			})
	}

	res := ProductReportResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
