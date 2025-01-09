package products_controller

import (
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) GetReportRecords(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var reportRecords []models.ProductReportRecords
	var err error

	if id == "" {
		reportRecords, err = p.SV.GetReportRecords(0)
	} else {
		productId, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		reportRecords, err = p.SV.GetReportRecords(productId)
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data []ReportRecordFullJSON
	for _, r := range reportRecords {
		data = append(data,
			ReportRecordFullJSON{
				ProductId:    r.ProductId,
				Description:  r.Description,
				RecordsCount: r.RecordsCount,
			})
	}

	res := ReportRecordsResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
