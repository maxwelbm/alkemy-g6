package warehouses_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.Service.GetById(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	err = c.Service.Delete(id)
	if errors.Is(err, service.ErrWarehouseServiceEmployeesAssociated) || errors.Is(err, service.ErrWarehouseServiceSectionsAssociated) {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotDeleteOrUpdateParentRow {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	res := WarehouseResJSON{
		Message: "Success",
		Data:    nil,
	}
	response.JSON(w, http.StatusNoContent, res)
}
