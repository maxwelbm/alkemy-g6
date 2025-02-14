package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type InboundOrderFactory struct {
	db *sql.DB
}

func NewInboundOrderFactory(db *sql.DB) *InboundOrderFactory {
	return &InboundOrderFactory{db: db}
}

func defaultInboundOrder() models.InboundOrder {
	return models.InboundOrder{
		OrderDate:      randstr.Alphanumeric(8),
		OrderNumber:    1,
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}
}

func (f *InboundOrderFactory) Create(inboundOrder models.InboundOrder) (record models.InboundOrder, err error) {
	populateInboundOrderParams(&inboundOrder)

	if err = f.checkEmployeeExists(inboundOrder.EmployeeID); err != nil {
		return inboundOrder, err
	}

	if err = f.checkProductBatchExists(inboundOrder.ProductBatchID); err != nil {
		return inboundOrder, err
	}

	if err = f.checkWarehouseExists(inboundOrder.WarehouseID); err != nil {
		return inboundOrder, err
	}

	query := `
		INSERT INTO inbound_orders 
			(
			%s
			order_date,
			order_number,
			employee_id,
			product_batch_id,
			warehouse_id
			) 
		VALUES (%s?, ?, ?, ?, ?)
	`

	switch inboundOrder.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(inboundOrder.ID)+",")
	}

	_, err = f.db.Exec(query,
		inboundOrder.OrderDate,
		inboundOrder.OrderNumber,
		inboundOrder.EmployeeID,
		inboundOrder.ProductBatchID,
		inboundOrder.WarehouseID,
	)

	return inboundOrder, err
}

func populateInboundOrderParams(inboundOrder *models.InboundOrder) {
	defaultInboundOrder := defaultInboundOrder()
	if inboundOrder == nil {
		inboundOrder = &defaultInboundOrder
	}

	if inboundOrder.OrderDate == "" {
		inboundOrder.OrderDate = defaultInboundOrder.OrderDate
	}

	if inboundOrder.OrderNumber == 0 {
		inboundOrder.OrderNumber = defaultInboundOrder.OrderNumber
	}

	if inboundOrder.EmployeeID == 0 {
		inboundOrder.EmployeeID = defaultInboundOrder.EmployeeID
	}

	if inboundOrder.ProductBatchID == 0 {
		inboundOrder.ProductBatchID = defaultInboundOrder.ProductBatchID
	}

	if inboundOrder.WarehouseID == 0 {
		inboundOrder.WarehouseID = defaultInboundOrder.WarehouseID
	}
}

func (f *InboundOrderFactory) checkEmployeeExists(employeeID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM employees WHERE id = ?`, employeeID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createEmployee()

	return
}

func (f *InboundOrderFactory) checkProductBatchExists(productBatchID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM product_batches WHERE id = ?`, productBatchID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createProductBatch()

	return
}

func (f *InboundOrderFactory) checkWarehouseExists(warehouseID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM warehouses WHERE id = ?`, warehouseID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createWarehouse()

	return
}

func (f *InboundOrderFactory) createEmployee() (err error) {
	warehouseFactory := NewEmployeeFactory(f.db)
	_, err = warehouseFactory.Create(models.Employee{})

	return
}

func (f *InboundOrderFactory) createProductBatch() (err error) {
	warehouseFactory := NewProductBatchesFactory(f.db)
	_, err = warehouseFactory.Create(models.ProductBatches{})

	return
}

func (f *InboundOrderFactory) createWarehouse() (err error) {
	warehouseFactory := NewWarehouseFactory(f.db)
	_, err = warehouseFactory.Create(models.Warehouse{})

	return
}
