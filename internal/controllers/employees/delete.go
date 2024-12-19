package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/employees"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *Employees) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = c.sv.Delete(id)
	if errors.Is(err, repository.ErrEmployeesRepositoryNotFound) {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, nil)
		return
	}

	response.JSON(w, http.StatusNoContent, "Sucess delete")
}
