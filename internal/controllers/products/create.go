package products_controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (p *ProductsDefault) Create(w http.ResponseWriter, r *http.Request) {
	var prodJson NewProductAttributesJSON
	json.NewDecoder(r.Body).Decode(&prodJson)

	err := prodJson.validate()
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	prodDTO := models.ProductDTO{
		ProductCode:    *prodJson.ProductCode,
		Description:    *prodJson.Description,
		Height:         *prodJson.Height,
		Length:         *prodJson.Length,
		Width:          *prodJson.Width,
		Weight:         *prodJson.Weight,
		ExpirationRate: *prodJson.ExpirationRate,
		FreezingRate:   *prodJson.FreezingRate,
		RecomFreezTemp: *prodJson.RecomFreezTemp,
		ProductTypeID:  *prodJson.ProductTypeID,
	}
	if prodJson.SellerID != nil {
		prodDTO.SellerID = *prodJson.SellerID
	}

	newProd, err := p.SV.Create(prodDTO)
	if errors.Is(err, models.ErrProductUniqueness) {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}

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
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	res := ProductResJSON{Message: "Created", Data: newProd}
	response.JSON(w, http.StatusCreated, res)
}
