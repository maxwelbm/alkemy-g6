package sections_controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *SectionsController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid Id format ")
		return
	}

	_, err = c.SV.GetById(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	err = c.SV.Delete(id)
	if err != nil {
		// Check if the error is a MySQL foreign key constraint error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotDeleteOrUpdateParentRow {
			// If it is, return a conflict error
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, return an internal server error
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := SectionResJSON{
		Message: "No content",
	}
	response.JSON(w, http.StatusNoContent, res)
}
