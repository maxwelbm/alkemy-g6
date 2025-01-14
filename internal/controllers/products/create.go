package productsctl

import (
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Create handles the creation of a new product.
// @Summary Create a new product
// @Description Create a new product with the provided attributes
// @Tags products
// @Accept json
// @Produce json
// @Param product body NewProductAttributesJSON true "Product attributes"
// @Success 201 {object} ProductResJSON "Created"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/products [post]
func (p *ProductsDefault) Create(w http.ResponseWriter, r *http.Request) {
	var prodJSON NewProductAttributesJSON
	if err := json.NewDecoder(r.Body).Decode(&prodJSON); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err := prodJSON.validate()
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	prodDTO := models.ProductDTO{
		ProductCode:    *prodJSON.ProductCode,
		Description:    *prodJSON.Description,
		Height:         *prodJSON.Height,
		Length:         *prodJSON.Length,
		Width:          *prodJSON.Width,
		NetWeight:      *prodJSON.NetWeight,
		ExpirationRate: *prodJSON.ExpirationRate,
		FreezingRate:   *prodJSON.FreezingRate,
		RecomFreezTemp: *prodJSON.RecomFreezTemp,
		ProductTypeID:  *prodJSON.ProductTypeID,
		SellerID:       *prodJSON.SellerID,
	}

	newProd, err := p.SV.Create(prodDTO)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case mysqlerr.CodeDuplicateEntry:
				response.Error(w, http.StatusConflict, err.Error())
				return
			case mysqlerr.CodeCannotAddOrUpdateChildRow:
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	res := ProductResJSON{Message: "Created", Data: newProd}
	response.JSON(w, http.StatusCreated, res)
}
