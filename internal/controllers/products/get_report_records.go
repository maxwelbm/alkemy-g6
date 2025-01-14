package productsctl

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// GetReportRecords handles the HTTP request to retrieve report records for a specific product.
// @Summary Retrieve report records for a product
// @Description Retrieves report records for a product based on the provided product ID.
// @Tags products
// @Produce json
// @Param id query int true "Product ID"
// @Success 200 {object} ReportRecordsResJSON "Successfully retrieved report records"
// @Failure 400 {object} response.ErrorResponse "Invalid product ID or bad request"
// @Failure 404 {object} response.ErrorResponse "Product not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/products/reportRecords [get]
func (p *ProductsDefault) GetReportRecords(w http.ResponseWriter, r *http.Request) {
	var productID int

	var err error

	id := r.URL.Query().Get("id")
	if id != "" {
		productID, err = strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if productID < 1 {
			response.Error(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}

	var reportRecords []models.ProductReportRecords
	reportRecords, err = p.SV.GetReportRecords(productID)

	if err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	data := make([]ReportRecordFullJSON, 0, len(reportRecords))
	for _, r := range reportRecords {
		data = append(data,
			ReportRecordFullJSON{
				ProductID:    r.ProductID,
				Description:  r.Description,
				RecordsCount: r.RecordsCount,
			})
	}

	res := ReportRecordsResJSON{Data: data}
	response.JSON(w, http.StatusOK, res)
}
