package sellerController

import (
	"net/http"

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
