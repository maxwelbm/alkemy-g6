package inboundordersctl

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
func (c *InboundOrdersController) Create(w http.ResponseWriter, r *http.Request) {
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
		EmployeeId:     inboundOrdersJSON.EmployeeId,
		ProductBatchId: inboundOrdersJSON.ProductBatchId,
		WarehouseId:    inboundOrdersJSON.WarehouseId,
	}

	inb, err := c.SV.Create(inboundOrders)
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
		Message: "Sucess created",
		Data: InboundOrdersAttributes{
			ID:             inb.ID,
			OrderDate:      inb.OrderDate,
			OrderNumber:    inb.OrderNumber,
			EmployeeId:     inb.EmployeeId,
			ProductBatchId: inb.ProductBatchId,
			WarehouseId:    inb.WarehouseId,
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

	if inboundOrders.EmployeeId == nil || *inboundOrders.EmployeeId <= 0 {
		errosInbound = append(errosInbound, "error: attribute Employee Id invalid")
	}

	if inboundOrders.ProductBatchId == nil || *inboundOrders.ProductBatchId <= 0 {
		errosInbound = append(errosInbound, "error: attribute Product Batch Id invalid")
	}

	if inboundOrders.WarehouseId == nil || *inboundOrders.WarehouseId <= 0 {
		errosInbound = append(errosInbound, "error: attribute Warehouse Id invalid")
	}

	if len(errosInbound) > 0 {
		err = errors.New(fmt.Sprintf("validation errors: %v", errosInbound))
	}
	return
}
