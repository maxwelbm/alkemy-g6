package sellerController

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (controller *SellerDefault) PatchSeller(w http.ResponseWriter, r *http.Request) {
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

	var sellerRequest SellerRequestPatch
	if err := json.NewDecoder(r.Body).Decode(&sellerRequest); err != nil {
		response.JSON(w, http.StatusBadRequest, "Error ao decodificarJSON")
		return
	}

	if sellerRequest.ID == nil {
		sellerRequest.ID = &seller.ID
	}
	if sellerRequest.CID == nil {
		sellerRequest.CID = &seller.CID
	}
	if sellerRequest.CompanyName == nil {
		sellerRequest.CompanyName = &seller.CompanyName
	}
	if sellerRequest.Address == nil {
		sellerRequest.Address = &seller.Address
	}
	if sellerRequest.Telephone == nil {
		sellerRequest.Telephone = &seller.Telephone
	}

	sellerToUpdate := modelsSeller.Seller{
		ID:          id,
		CID:         *sellerRequest.CID,
		CompanyName: *sellerRequest.CompanyName,
		Address:     *sellerRequest.Address,
		Telephone:   *sellerRequest.Telephone,
	}

	errToPatch := controller.sv.PatchSeller(sellerToUpdate)

	if errToPatch != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to done the patch")
		return
	}

	data := SellerDataResJSON{
		ID:          sellerToUpdate.ID,
		CID:         sellerToUpdate.CID,
		CompanyName: sellerToUpdate.CompanyName,
		Address:     sellerToUpdate.Address,
		Telephone:   sellerToUpdate.Telephone,
	}

	res := SellerResJSON{
		Message: "Success",
		Data:    data,
	}
	response.JSON(w, http.StatusOK, res)

}
