package buyers_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (ct *BuyersDefault) ReportPurchaseOrders(w http.ResponseWriter, r *http.Request) {
	var id int
	var err error

	param := r.URL.Query().Get("id")
	if param != "" {
		id, err = strconv.Atoi(param)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		if id < 1 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}
	list, err := ct.SV.ReportPurchaseOrders(id)
	if err != nil {
		if errors.Is(err, models.ErrBuyerNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data []models.BuyerPurchaseOrdersReportJSON
	for _, value := range list {
		result := models.BuyerPurchaseOrdersReportJSON{
			Id:                  value.ID,
			CardNumberId:        value.CardNumberId,
			FirstName:           value.FirstName,
			LastName:            value.LastName,
			PurchaseOrdersCount: value.PurchaseOrdersCount,
		}
		data = append(data, result)
	}

	res := BuyerResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
