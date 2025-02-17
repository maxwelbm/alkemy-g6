package inboundordersrp

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestInboundOrdersRepository_Create(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewInboundOrderFactory(db)
	inboundOrders := factory.Build(models.InboundOrder{ID: 1})

	type arg struct {
		dto models.InboundOrderDTO
	}
	type expected struct {
		inboundOrders models.InboundOrder
		err           error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		expected
	}{
		{
			name: "When successfully creating a new inbound order",
			setup: func() {
				_, err := factories.NewEmployeeFactory(db).Create(models.Employee{ID: 2})
				require.NoError(t, err)
				_, err = factories.NewProductBatchesFactory(db).Create(models.ProductBatches{ID: 1})
				require.NoError(t, err)
				_, err = factories.NewWarehouseFactory(db).Create(models.Warehouse{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.InboundOrderDTO{
					OrderDate:      &inboundOrders.OrderDate,
					OrderNumber:    &inboundOrders.OrderNumber,
					EmployeeID:     &inboundOrders.EmployeeID,
					ProductBatchID: &inboundOrders.ProductBatchID,
					WarehouseID:    &inboundOrders.WarehouseID,
				},
			},
			expected: expected{
				inboundOrders: inboundOrders,
				err:           nil,
			},
		},
		{
			name: "Error - When creating a Inbound Order from a non-existent Employee",
			arg: arg{
				dto: models.InboundOrderDTO{
					OrderDate:      &inboundOrders.OrderDate,
					OrderNumber:    &inboundOrders.OrderNumber,
					EmployeeID:     &inboundOrders.EmployeeID,
					ProductBatchID: &inboundOrders.ProductBatchID,
					WarehouseID:    &inboundOrders.WarehouseID,
				},
			},
			expected: expected{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			},
		},
		{
			name: "Error - When creating a Inbound Order with a duplicated Order Number",
			setup: func() {
				_, err := factory.Create(inboundOrders)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.InboundOrderDTO{
					OrderDate:      &inboundOrders.OrderDate,
					OrderNumber:    &inboundOrders.OrderNumber,
					EmployeeID:     &inboundOrders.EmployeeID,
					ProductBatchID: &inboundOrders.ProductBatchID,
					WarehouseID:    &inboundOrders.WarehouseID,
				},
			},
			expected: expected{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			// Arrange
			rp := NewInboundOrdersRepository(db)
			// Act
			got, err := rp.Create(tt.dto)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.expected.inboundOrders, got)
		})
	}
}
