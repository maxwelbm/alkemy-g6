package buyerController

import (
	"encoding/json"
	"net/http"

	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (controller *BuyerDefault) PostBuyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var buyerRequest BuyerRequestPost
	if err := json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		response.JSON(w, http.StatusBadRequest, "Error ao decodificarJSON")
		return
	}

	if buyerRequest.CardNumberId == nil {
		response.JSON(w, http.StatusUnprocessableEntity, "CardNumberId inexists in requests!")
		return
	}

	_, errCardNumberId := controller.sv.GetByCardNumberId(*buyerRequest.CardNumberId)

	if errCardNumberId == nil {
		response.Error(w, http.StatusConflict, "Card Number Id already exists!")
		return
	}

	buyerToCreate := modelsBuyer.Buyer{
		CardNumberId: *buyerRequest.CardNumberId,
		FirstName:    *buyerRequest.FirstName,
		LastName:     *buyerRequest.LastName,
	}

	buyerCreated, errToPost := controller.sv.PostBuyer(buyerToCreate)

	if errToPost != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to done the post")
		return
	}

	data := BuyerDataResJSON{
		Id:           buyerCreated.Id,
		CardNumberId: buyerCreated.CardNumberId,
		FirstName:    buyerCreated.FirstName,
		LastName:     buyerCreated.LastName,
	}

	res := BuyerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusCreated, res)

}
