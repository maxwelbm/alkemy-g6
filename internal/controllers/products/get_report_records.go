package products_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) GetReportRecords(w http.ResponseWriter, r *http.Request) {
	var productId int
	var err error

	id := r.URL.Query().Get("id")
	if id != "" {
		productId, err = strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		if productId < 1 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}	
	
	var reportRecords []models.ProductReportRecords
	reportRecords, err = p.SV.GetReportRecords(productId)

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
