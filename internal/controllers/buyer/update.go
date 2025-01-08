package buyers_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (ct *BuyersController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, "Failed to convert request id")
		return
	}

	var buyerRequest BuyerRequestPatch
	if err := json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		response.JSON(w, http.StatusBadRequest, "Error ao decodificarJSON")
		return
	}

	buyerToUpdate := models.BuyerDTO{
		Id:           &id,
		CardNumberId: buyerRequest.CardNumberId,
		FirstName:    buyerRequest.FirstName,
		LastName:     buyerRequest.LastName,
	}

	buyerReturn, err := ct.SV.Update(buyerToUpdate)

	if errors.Is(err, models.ErrorIdNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal error! Failed to done the patch")
		return
	}

	data := BuyerDataResJSON{
		Id:           buyerReturn.Id,
		CardNumberId: buyerReturn.CardNumberId,
		FirstName:    buyerReturn.FirstName,
		LastName:     buyerReturn.LastName,
	}

	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)

}
