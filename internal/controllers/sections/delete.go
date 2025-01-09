package sections_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid Id format ")
		return
	}

	err = c.SV.Delete(id)

	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	// if errors.Is(err, repository.ErrSectionNotFound) {
	// 	response.Error(w, http.StatusNotFound, err.Error())
	// 	return
	// }
	// if err != nil {
	// 	response.Error(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	res := SectionResJSON{
		Message: "No content",
	}
	response.JSON(w, http.StatusNoContent, res)
}
