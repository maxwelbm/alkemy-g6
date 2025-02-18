package inboundordersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (e *InboundOrdersRepository) Create(inboundOrders models.InboundOrderDTO) (newInboundOrders models.InboundOrder, err error) {
	query := `
		INSERT INTO inbound_orders
			(order_date,order_number, employee_id, product_batch_id, warehouse_id) 
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := e.db.Exec(
		query,
		*inboundOrders.OrderDate,
		*inboundOrders.OrderNumber,
		*inboundOrders.EmployeeID,
		*inboundOrders.ProductBatchID,
		*inboundOrders.WarehouseID,
	)

	if err != nil {
		return newInboundOrders, err
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return newInboundOrders, err
	}

	query = "SELECT id, order_date, order_number, employee_id, product_batch_id, warehouse_id FROM inbound_orders WHERE id = ?"
	err = e.db.
		QueryRow(query, lastInsertID).
		Scan(
			&newInboundOrders.ID,
			&newInboundOrders.OrderDate,
			&newInboundOrders.OrderNumber,
			&newInboundOrders.EmployeeID,
			&newInboundOrders.ProductBatchID,
			&newInboundOrders.WarehouseID,
		)

	return newInboundOrders, err
}
