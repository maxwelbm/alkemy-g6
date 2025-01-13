package purchase_orders_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (pc *PurchaseOrdersController) Create(w http.ResponseWriter, r *http.Request) {
	var purchaseOrdersJson PurchaseOrdersJSON
	err := json.NewDecoder(r.Body).Decode(&purchaseOrdersJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = purchaseOrdersJson.validate(); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	poDTO := &models.PurchaseOrdersDTO{
		OrderNumber:     *purchaseOrdersJson.OrderNumber,
		OrderDate:       *purchaseOrdersJson.OrderDate,
		TrackingCode:    *purchaseOrdersJson.TrackingCode,
		BuyerID:         *purchaseOrdersJson.BuyerId,
		ProductRecordID: *purchaseOrdersJson.ProductRecordId,
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
		BuyerId:         purchaseOrders.BuyerID,
		ProductRecordId: purchaseOrders.ProductRecordID,
	}

	res := ResPurchaseOrdersJSON{
		Message: "Success",
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
	if p.BuyerId == nil {
		poErrs = append(poErrs, "error: buyer_id: is required")
	}
	if p.ProductRecordId == nil {
		poErrs = append(poErrs, "error: product_record_id is required")
	}
	if len(poErrs) > 0 {
		err = errors.New(fmt.Sprintf("validation errors: %v", poErrs))
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
