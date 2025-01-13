package sections_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

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
