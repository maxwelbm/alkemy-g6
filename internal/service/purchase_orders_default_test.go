package service_test

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	purchaseordersrp "github.com/maxwelbm/alkemy-g6/internal/repository/purchase_orders"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

var purchaseOrders = []models.PurchaseOrders{
	{
		ID:              1,
		OrderNumber:     "O001",
		OrderDate:       "2021-09-01",
		TrackingCode:    "T001",
		BuyerID:         1,
		ProductRecordID: 1,
	},
	{
		ID:              2,
		OrderNumber:     "O002",
		OrderDate:       "2021-09-02",
		TrackingCode:    "T002",
		BuyerID:         2,
		ProductRecordID: 2,
	},
	{
		ID:              3,
		OrderNumber:     "O003",
		OrderDate:       "2021-09-03",
		TrackingCode:    "T003",
		BuyerID:         3,
		ProductRecordID: 3,
	},
}

func TestPurchaseOrdersDefault_Create(t *testing.T) {
	test := []struct {
		name                   string
		purchaseOrders         models.PurchaseOrders
		err                    error
		expectedPurchaseOrders models.PurchaseOrders
		expectedErr            error
	}{
		{
			name:                   "When the repository returns a purchase orders",
			purchaseOrders:         purchaseOrders[0],
			err:                    nil,
			expectedPurchaseOrders: purchaseOrders[0],
			expectedErr:            nil,
		},
		{
			name:                   "When the repository returns an error",
			purchaseOrders:         purchaseOrders[0],
			err:                    &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			expectedPurchaseOrders: purchaseOrders[0],
			expectedErr:            &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := purchaseordersrp.NewPurchaseOrdersRepositoryMock()
			sv := service.NewPurchaseOrdersService(rp)
			dto := models.PurchaseOrdersDTO{
				OrderNumber:     tt.purchaseOrders.OrderNumber,
				OrderDate:       tt.purchaseOrders.OrderDate,
				TrackingCode:    tt.purchaseOrders.TrackingCode,
				BuyerID:         tt.purchaseOrders.BuyerID,
				ProductRecordID: tt.purchaseOrders.ProductRecordID,
			}

			// Act
			rp.On("Create", dto).Return(tt.purchaseOrders, tt.err)
			result, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.expectedPurchaseOrders, result)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
