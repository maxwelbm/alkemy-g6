package repository

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

func (e *InboundOrdersRepository) Create(inboundOrders models.InboundOrdersDTO) (newInboundOrders models.InboundOrders, err error) {
	query := "INSERT INTO inbound_orders (order_date,order_number, employee_id, product_batch_id, warehouse_id) VALUES (?, ?, ?, ?, ?)"

	result, err := e.db.Exec(query, inboundOrders.OrderDate, inboundOrders.OrderNumber, inboundOrders.EmployeeId, inboundOrders.ProductBatchId, inboundOrders.WarehouseId)
	if err != nil {
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return
	}

	query = "SELECT id, order_date, order_number, employee_id, product_batch_id, warehouse_id FROM inbound_orders WHERE id = ?"
	err = e.db.
		QueryRow(query, lastInsertID).
		Scan(
			&newInboundOrders.ID,
			&newInboundOrders.OrderDate,
			&newInboundOrders.OrderNumber,
			&newInboundOrders.EmployeeId,
			&newInboundOrders.ProductBatchId,
			&newInboundOrders.WarehouseId,
		)

	return
}
