package sellerController

import (
	"encoding/json"
	"net/http"
)

func (controller *SellerController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		sellers, err := controller.service.FindAll()

		if err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(SellerResponse{Status: http.StatusOK, Message: "Get realizado com sucesso!", Data: nil})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SellerResponse{Status: http.StatusOK, Message: "Get realizado com sucesso!", Data: sellers})

	}
}
