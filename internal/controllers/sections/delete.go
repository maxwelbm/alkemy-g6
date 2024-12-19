package sections

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsDefault) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid Id format ")
		return
	}

	err = c.sv.Delete(id)

	if err != nil {
		response.JSON(w, http.StatusNotFound, err.Error())
		return
	}

	res := SectionResJSON{
		Message: "No content",
	}
	response.JSON(w, http.StatusNoContent, res)
}
