package inboundOrders_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/response"
)

func (c *InboundOrdersController) Create(w http.ResponseWriter, r *http.Request) {
	var inboundOrdersJson InboundOrdersReqJSON
	err := json.NewDecoder(r.Body).Decode(&inboundOrdersJson)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Body invalid")
		return
	}

	err = validateNewInboundOrders(inboundOrdersJson)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	inboundOrders := models.InboundOrdersDTO{
		OrderDate:      inboundOrdersJson.OrderDate,
		OrderNumber:    inboundOrdersJson.OrderNumber,
		EmployeeId:     inboundOrdersJson.EmployeeId,
		ProductBatchId: inboundOrdersJson.ProductBatchId,
		WarehouseId:    inboundOrdersJson.WarehouseId,
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
