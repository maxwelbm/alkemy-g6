package purchaseordersctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Create handles the HTTP request for creating a new purchase order.
// @Summary Create a new purchase order
// @Description Accepts a JSON payload, validates it, and creates a purchase order in the system.
// @Tags purchase_orders
// @Accept json
// @Produce json
// @Param purchaseOrder body PurchaseOrdersJSON true "Purchase Order JSON"
// @Success 201 {object} ResPurchaseOrdersJSON "Created"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router 	/api/v1/purchaseOrders [post]
func (pc *PurchaseOrdersController) Create(w http.ResponseWriter, r *http.Request) {
	var purchaseOrdersJSON PurchaseOrdersJSON

	err := json.NewDecoder(r.Body).Decode(&purchaseOrdersJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = purchaseOrdersJSON.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	poDTO := &models.PurchaseOrdersDTO{
		OrderNumber:     *purchaseOrdersJSON.OrderNumber,
		OrderDate:       *purchaseOrdersJSON.OrderDate,
		TrackingCode:    *purchaseOrdersJSON.TrackingCode,
		BuyerID:         *purchaseOrdersJSON.BuyerID,
		ProductRecordID: *purchaseOrdersJSON.ProductRecordID,
	}

	purchaseOrders, err := pc.sv.Create(*poDTO)
	if err != nil {
		pc.handleCreateError(w, err)
		return
	}

	data := PurchaseOrdersResJSON{
		ID:              purchaseOrders.ID,
		OrderNumber:     purchaseOrders.OrderNumber,
		OrderDate:       purchaseOrders.OrderDate,
		TrackingCode:    purchaseOrders.TrackingCode,
		BuyerID:         purchaseOrders.BuyerID,
		ProductRecordID: purchaseOrders.ProductRecordID,
	}

	res := ResPurchaseOrdersJSON{
		Message: http.StatusText(http.StatusCreated),
		Data:    data,
	}
	response.JSON(w, http.StatusCreated, res)
}

func (p *PurchaseOrdersJSON) validate() (err error) {
	var poErrs []string

	if p.OrderNumber == nil {
		poErrs = append(poErrs, "error: order_number is required")
	}

	if p.OrderDate == nil {
		poErrs = append(poErrs, "error: order_date is required")
	}

	if p.TrackingCode == nil {
		poErrs = append(poErrs, "error: tracking_code is required")
	}

	if p.BuyerID == nil {
		poErrs = append(poErrs, "error: buyer_id: is required")
	}

	if p.ProductRecordID == nil {
		poErrs = append(poErrs, "error: product_record_id is required")
	}

	if len(poErrs) > 0 {
		err = fmt.Errorf("validation errors: %v", poErrs)
	}

	return
}

func (pc *PurchaseOrdersController) handleCreateError(w http.ResponseWriter, err error) {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case mysqlerr.CodeDuplicateEntry:
			response.Error(w, http.StatusConflict, err.Error())
		case mysqlerr.CodeIncorrectDateValue:
			response.Error(w, http.StatusBadRequest, err.Error())
		case mysqlerr.CodeCannotAddOrUpdateChildRow:
			response.Error(w, http.StatusConflict, err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.Error(w, http.StatusInternalServerError, err.Error())
}
