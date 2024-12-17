package loaders

import (
	"encoding/json"
	"os"

	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
)

type ProductJSONFile struct {
	path string
}

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{
		path: path,
	}
}

type ProductJSON struct {
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

func (l *ProductJSONFile) Load() (products map[int]models.Product, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var productsJSON []ProductJSON
	err = json.NewDecoder(file).Decode(&productsJSON)
	if err != nil {
		return
	}

	// serialize products
	products = make(map[int]models.Product)
	for _, s := range productsJSON {
		products[s.ID] = models.Product{
			ID:             s.ID,
			ProductCode:    s.ProductCode,
			Description:    s.Description,
			Height:         s.Height,
			Length:         s.Length,
			Width:          s.Width,
			Weight:         s.Weight,
			ExpirationRate: s.ExpirationRate,
			FreezingRate:   s.FreezingRate,
			RecomFreezTemp: s.RecomFreezTemp,
			ProductTypeID:  s.ProductTypeID,
			SellerID:       s.SellerID,
		}
	}

	return
}
