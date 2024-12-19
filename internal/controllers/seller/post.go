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

	if sellerRequest.ID == nil || sellerRequest.CID == nil {
		response.JSON(w, http.StatusUnprocessableEntity, "Id or Cid inexists in requests!")
		return
	}

	_, errId := controller.sv.GetById(*sellerRequest.ID)

	if errId == nil {
		response.Error(w, http.StatusConflict, "Id already exists!")
		return
	}

	_, errCid := controller.sv.GetByCid(*sellerRequest.CID)

	if errCid == nil {
		response.Error(w, http.StatusConflict, "Cid already exists!")
		return
	}

	sellerToCreate := modelsSeller.Seller{
		ID:          *sellerRequest.ID,
		CID:         *sellerRequest.CID,
		CompanyName: *sellerRequest.CompanyName,
		Address:     *sellerRequest.Address,
		Telephone:   *sellerRequest.Telephone,
	}

	errToPost := controller.sv.PostSeller(sellerToCreate)

	if errToPost != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to done the post")
		return
	}

	data := SellerDataResJSON{
		ID:          sellerToCreate.ID,
		CID:         sellerToCreate.CID,
		CompanyName: sellerToCreate.CompanyName,
		Address:     sellerToCreate.Address,
		Telephone:   sellerToCreate.Telephone,
	}

	res := SellerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusCreated, res)

}
