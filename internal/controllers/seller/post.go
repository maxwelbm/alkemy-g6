package sellerController

import (
	"encoding/json"
	"net/http"

	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (controller *SellerDefault) PostSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var sellerRequest SellerRequestPost
	if err := json.NewDecoder(r.Body).Decode(&sellerRequest); err != nil {
		response.JSON(w, http.StatusBadRequest, "Error ao decodificarJSON")
		return
	}

	if sellerRequest.CID == nil {
		response.JSON(w, http.StatusUnprocessableEntity, "Cid inexists in requests!")
		return
	}

	_, errCid := controller.sv.GetByCid(*sellerRequest.CID)

	if errCid == nil {
		response.Error(w, http.StatusConflict, "Cid already exists!")
		return
	}

	sellerToCreate := modelsSeller.Seller{
		ID:          0,
		CID:         *sellerRequest.CID,
		CompanyName: *sellerRequest.CompanyName,
		Address:     *sellerRequest.Address,
		Telephone:   *sellerRequest.Telephone,
	}

	sellerCreated, errToPost := controller.sv.PostSeller(sellerToCreate)

	if errToPost != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to done the post")
		return
	}

	data := SellerDataResJSON{
		ID:          sellerCreated.ID,
		CID:         sellerCreated.CID,
		CompanyName: sellerCreated.CompanyName,
		Address:     sellerCreated.Address,
		Telephone:   sellerCreated.Telephone,
	}

	res := SellerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusCreated, res)

}
