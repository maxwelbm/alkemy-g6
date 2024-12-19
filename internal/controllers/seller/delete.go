package sellerController

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (controller *SellerDefault) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, "Failed to convert request id")
		return
	}

	_, err = controller.sv.GetById(id)

	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	controller.sv.Delete(id)

	response.JSON(w, http.StatusNoContent, nil)

}
