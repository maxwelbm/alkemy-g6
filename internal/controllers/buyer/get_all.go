package buyers_controller

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (ct *BuyersDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	buyers, err := ct.SV.GetAll()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data []BuyerDataResJSON
	for _, value := range buyers {
		new := BuyerDataResJSON{
			Id:           value.Id,
			CardNumberId: value.CardNumberId,
			FirstName:    value.FirstName,
			LastName:     value.LastName,
		}

		data = append(data, new)
	}
	res := BuyersResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)

}
