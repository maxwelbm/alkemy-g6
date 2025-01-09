package buyers_controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (ct *BuyersDefault) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var buyerRequest BuyerRequestPost
	if err := json.NewDecoder(r.Body).Decode(&buyerRequest); err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	buyerToCreate := models.BuyerDTO{
		CardNumberId: buyerRequest.CardNumberId,
		FirstName:    buyerRequest.FirstName,
		LastName:     buyerRequest.LastName,
	}

	buyerCreated, err := ct.SV.Create(buyerToCreate)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
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
