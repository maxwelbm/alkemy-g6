package sellerController

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (controller *SellerDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	sellers, err := controller.sv.GetAll()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve sellers")
		return
	}

	var data []SellerDataResJSON
	for _, value := range sellers {
		new := SellerDataResJSON{
			ID:          value.ID,
			CID:         value.CID,
			CompanyName: value.CompanyName,
			Address:     value.Address,
			Telephone:   value.Telephone,
		}

		data = append(data, new)
	}
	res := SellersResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)

}

func (controller *SellerDefault) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, "Failed to convert request id")
		return
	}

	seller, err := controller.sv.GetById(id)

	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	var data SellerDataResJSON
	data = SellerDataResJSON{
		ID:          seller.ID,
		CID:         seller.CID,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
	}

	res := SellerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)

}
