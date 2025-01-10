package products_controller

import (
	"errors"
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
		productId, convErr := strconv.Atoi(id)
		if convErr != nil {
			response.Error(w, http.StatusInternalServerError, convErr.Error())
			return
		}

		reportRecords, err = p.SV.GetReportRecords(productId)
	}

	if err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
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
