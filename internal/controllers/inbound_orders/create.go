package inboundordersctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

// Create handles the creation of a new inbound order.
// @Summary Create a new inbound order
// @Description Create a new inbound order with the provided JSON data
// @Tags inbound_orders
// @Accept json
// @Produce json
// @Param inboundOrder body InboundOrdersReqJSON true "Inbound Order JSON"
// @Success 201 {object} InboundOrdersResJSON "Created"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 422 {object} response.ErrorResponse "Unprocessable Entity"
// @Failure 409 {object} response.ErrorResponse "Conflict"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/v1/inboundOrders [post]
func (ctl *InboundOrdersController) Create(w http.ResponseWriter, r *http.Request) {
	var inboundOrdersJSON InboundOrdersReqJSON
	err := json.NewDecoder(r.Body).Decode(&inboundOrdersJSON)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateNewInboundOrders(inboundOrdersJSON)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	inboundOrders := models.InboundOrdersDTO{
		OrderDate:      inboundOrdersJSON.OrderDate,
		OrderNumber:    inboundOrdersJSON.OrderNumber,
		EmployeeID:     inboundOrdersJSON.EmployeeID,
		ProductBatchID: inboundOrdersJSON.ProductBatchID,
		WarehouseID:    inboundOrdersJSON.WarehouseID,
	}

	inb, err := ctl.SV.Create(inboundOrders)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == mysqlerr.CodeCannotAddOrUpdateChildRow || mysqlErr.Number == mysqlerr.CodeDuplicateEntry {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		// For any other error, respond with an internal server error status
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	data := InboundOrdersResJSON{
		Message: "Success",
		Data: InboundOrdersAttributes{
			ID:             inb.ID,
			OrderDate:      inb.OrderDate,
			OrderNumber:    inb.OrderNumber,
			EmployeeID:     inb.EmployeeID,
			ProductBatchID: inb.ProductBatchID,
			WarehouseID:    inb.WarehouseID,
		},
	}
	response.JSON(w, http.StatusCreated, data)
}

func validateNewInboundOrders(inboundOrders InboundOrdersReqJSON) (err error) {
	var errosInbound []string

	if inboundOrders.OrderDate == nil || *inboundOrders.OrderDate == "" {
		errosInbound = append(errosInbound, "error: attribute Order Date invalid")
	}

	if inboundOrders.OrderNumber == nil || *inboundOrders.OrderNumber <= 0 {
		errosInbound = append(errosInbound, "error: attribute Order Number invalid")
	}

	if inboundOrders.EmployeeID == nil || *inboundOrders.EmployeeID <= 0 {
		errosInbound = append(errosInbound, "error: attribute Employee ID invalid")
	}

	if inboundOrders.ProductBatchID == nil || *inboundOrders.ProductBatchID <= 0 {
		errosInbound = append(errosInbound, "error: attribute Product Batch ID invalid")
	}

	if inboundOrders.WarehouseID == nil || *inboundOrders.WarehouseID <= 0 {
		errosInbound = append(errosInbound, "error: attribute Warehouse ID invalid")
	}

	if len(errosInbound) > 0 {
		err = fmt.Errorf("validation errors: %v", errosInbound)
	}

	return
}
