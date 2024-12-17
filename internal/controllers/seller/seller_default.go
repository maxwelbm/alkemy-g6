package sellerController

import (
	"net/http"

	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
)

type SellerResponse struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message,omitempty"`
	Data    map[int]modelsSeller.Seller `json:"data,omitempty"`
}

type SellerController struct {
	service modelsSeller.SellerService
}

type SellerDefault struct {
	sv modelsSeller.SellerService
}

func NewSellerController(sellerService modelsSeller.SellerService) *SellerDefault {
	return &SellerDefault{sv: sellerService}
}

func (controller *SellerDefault) FindAll() (w http.ResponseWriter, r *http.Request) {
	return controller.FindAll()
}
