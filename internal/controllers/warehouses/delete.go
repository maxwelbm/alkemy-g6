package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/internal/models/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *WarehouseDefault) Delete(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        response.Error(w, http.StatusBadRequest, err.Error())
        return
    }
    err = c.service.Delete(id)
    if errors.Is(err, models.ErrWarehouseRepositoryNotFound) {
        response.Error(w, http.StatusNotFound, err.Error())
        return
    }
    if errors.Is(err, service.ErrWarehouseServiceEmployeesAssociated) || errors.Is(err, service.ErrWarehouseServiceSectionsAssociated) {
        response.Error(w, http.StatusConflict, err.Error())
        return
    }
    if err != nil {
        response.JSON(w, http.StatusInternalServerError, nil)
        return
    }

    res := WarehouseResJSON{
        Message: "Success",
        Data:    nil,
    }
    response.JSON(w, http.StatusNoContent, res)
}
