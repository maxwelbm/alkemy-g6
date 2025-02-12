package inboundordersrp

import (
	"database/sql"
	"log"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func inboundOrdersSeeds(db *sql.DB) {
	// Create dependencies
	_, err := db.Exec("INSERT INTO warehouses (id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (1, '123 Main St', '555-1234', 'WH001', 100, -10.0)")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("INSERT INTO localities (locality_name, province_name, country_name) VALUES ('Locality1', 'Province1', 'Country1');")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (1, 'Company1', '123 Main St', '555-1234', 1);")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("INSERT INTO products (product_code, description, height, length, width, net_weight, expiration_rate, freezing_rate, recommended_freezing_temperature, product_type_id, seller_id) VALUES ('P001', 'Product 1', 10.0, 20.0, 30.0, 40.0, 0.1, 0.2, -10.0, 1, 1);")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("INSERT INTO employees (id, card_number_id, first_name, last_name, warehouse_id) VALUES (1, 'EMP001', 'John', 'Doe', 1);")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("INSERT INTO sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (1, 'SC-001', -10.0, -20.0, 50, 10, 100, 1, 1);")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("INSERT INTO product_batches (id, batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature, due_date, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES (1, '1', 200, 100, -10, -20, '2023-12-31', '2023-01-01', '08:00:00', 1, 1)")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create inbound orders
	_, err = db.Exec("INSERT INTO inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (1, '2023-01-01', 1001, 1, 1, 1);")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestCreate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	type call struct {
		dto models.InboundOrdersDTO
	}
	type want struct {
		order models.InboundOrders
		wErr  error
	}
	tests := []struct {
		name  string
		setup func()
		call
		want
	}{
		{
			name: "When creating a valid inbound order",
			setup: func() {
				inboundOrdersSeeds(db)
			},
			call: call{
				dto: func() models.InboundOrdersDTO {
					orderDate := "2023-01-01"
					orderNumber := 2
					employeeID := 1
					productBatchID := 1
					warehouseID := 1

					io := models.InboundOrdersDTO{}
					io.OrderDate = &orderDate
					io.OrderNumber = &orderNumber
					io.EmployeeID = &employeeID
					io.ProductBatchID = &productBatchID
					io.WarehouseID = &warehouseID

					return io
				}(),
			},
			want: want{
				order: models.InboundOrders{
					ID:             2,
					OrderDate:      "2023-01-01T00:00:00Z",
					OrderNumber:    2,
					EmployeeID:     1,
					ProductBatchID: 1,
					WarehouseID:    1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cleans up all tables before test execution and populates with new records
			truncate()
			tt.setup()

			// Arrange
			rp := NewInboundOrdersRepository(db)
			// Act
			got, err := rp.Create(tt.dto)
			log.Println(got, err)

			// Assert
			if tt.wErr != nil {
				require.Equal(t, tt.wErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.order, got)
			}
		})
	}
}
