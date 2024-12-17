package products_controller

import models "github.com/maxwelbm/alkemy-g6/internal/models/products"

type ProductsDefault struct {
	sv models.ProductService
}

func NewProductsDefault(sv models.ProductService) *ProductsDefault {
	return &ProductsDefault{sv: sv}
}

type ProductFullJSON struct {
	ID             int     `json:"id"`
	ProductCode    string  `json:"product_code"`
	Description    string  `json:"description"`
	Height         float64 `json:"height"`
	Length         float64 `json:"length"`
	Width          float64 `json:"width"`
	Weight         float64 `json:"weight"`
	ExpirationRate float64 `json:"expiration_rate"`
	FreezingRate   float64 `json:"freezing_rate"`
	RecomFreezTemp float64 `json:"recommended_freezing_temp"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}
