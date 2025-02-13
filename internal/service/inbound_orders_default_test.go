package service_test

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	inboundOrdersrp "github.com/maxwelbm/alkemy-g6/internal/repository/inbound_orders"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

var InboundOrdersFixture = []models.InboundOrders{
	{
		ID:             1,
		OrderDate:      "2025-01-01",
		OrderNumber:    100,
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	},
	{
		ID:             2,
		OrderDate:      "2025-01-02",
		OrderNumber:    101,
		EmployeeID:     2,
		ProductBatchID: 2,
		WarehouseID:    2,
	},
}

func TestInboundOrdersDefault_Create(t *testing.T) {
	tests := []struct {
		name                string
		inboundOrder        models.InboundOrders
		err                 error
		wantedInboundOrders models.InboundOrders
		wantedErr           error
	}{
		{
			name:                "When the repository returns a inboundOrder",
			inboundOrder:        InboundOrdersFixture[0],
			err:                 nil,
			wantedInboundOrders: InboundOrdersFixture[0],
			wantedErr:           nil,
		},
		{
			name:                "When the repository returns an error",
			inboundOrder:        InboundOrdersFixture[0],
			err:                 &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedInboundOrders: InboundOrdersFixture[0],
			wantedErr:           &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := inboundOrdersrp.NewInboundOrdersRepositoryMock()
			sv := service.NewInboundOrdersService(rp)

			dto := models.InboundOrdersDTO{
				ID:             &tt.inboundOrder.ID,
				OrderDate:      &tt.inboundOrder.OrderDate,
				OrderNumber:    &tt.inboundOrder.OrderNumber,
				EmployeeID:     &tt.inboundOrder.EmployeeID,
				ProductBatchID: &tt.inboundOrder.ProductBatchID,
				WarehouseID:    &tt.inboundOrder.WarehouseID,
			}
			// Act
			rp.On("Create", dto).Return(tt.inboundOrder, tt.err)
			inboundOrder, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedInboundOrders, inboundOrder)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
