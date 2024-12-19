package buyerController

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (controller *BuyerDefault) PatchBuyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, "Failed to convert request id")
		return
	}

	buyer, err := controller.sv.GetById(id)

	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	var buyerRequest BuyerRequestPatch
	if err := json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		response.JSON(w, http.StatusBadRequest, "Error ao decodificarJSON")
		return
	}

	if buyerRequest.Id == nil {
		buyerRequest.Id = &buyer.Id
	}
	if buyerRequest.CardNumberId == nil {
		buyerRequest.CardNumberId = &buyer.CardNumberId
	}
	if buyerRequest.FirstName == nil {
		buyerRequest.FirstName = &buyer.FirstName
	}
	if buyerRequest.LastName == nil {
		buyerRequest.LastName = &buyer.LastName
	}

	buyerToUpdate := modelsBuyer.Buyer{
		Id:           id,
		CardNumberId: *buyerRequest.CardNumberId,
		FirstName:    *buyerRequest.FirstName,
		LastName:     *&buyer.LastName,
	}

	errToPatch := controller.sv.PatchBuyer(buyerToUpdate)

	if errToPatch != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to done the patch")
		return
	}

	data := BuyerDataResJSON{
		Id:           buyerToUpdate.Id,
		CardNumberId: buyerToUpdate.CardNumberId,
		FirstName:    buyerToUpdate.FirstName,
		LastName:     buyerToUpdate.LastName,
	}

	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)

}
