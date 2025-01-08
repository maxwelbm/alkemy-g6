package buyers_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (ct *BuyersController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, "Failed to convert request id")
		return
	}

	buyer, err := ct.SV.GetById(id)

	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	var data BuyerDataResJSON
	data = BuyerDataResJSON{
		Id:           buyer.Id,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}

	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)

}
