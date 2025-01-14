package productsctl

import (
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetAll handles the HTTP request to retrieve all products.
// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags products
// @Produce json
// @Success 200 {object} ProductResJSON "List of products"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/products [get]
func (p *ProductsDefault) GetAll(w http.ResponseWriter, r *http.Request) {
	prods, err := p.SV.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := make([]ProductFullJSON, len(prods))
	for _, p := range prods {
		data = append(data,
			ProductFullJSON{
				ID:             p.ID,
				ProductCode:    p.ProductCode,
				Description:    p.Description,
				Height:         p.Height,
				Length:         p.Length,
				Width:          p.Width,
				NetWeight:      p.NetWeight,
				ExpirationRate: p.ExpirationRate,
				FreezingRate:   p.FreezingRate,
				RecomFreezTemp: p.RecomFreezTemp,
				ProductTypeID:  p.ProductTypeID,
				SellerID:       p.SellerID,
			})
	}

	res := ProductResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
