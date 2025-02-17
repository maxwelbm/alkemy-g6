package purchaseordersrp

import (
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewPurchaseOrdersFactory(db)
	fixture := factory.Build(models.PurchaseOrders{ID: 1})

	type arg struct {
		dto models.PurchaseOrdersDTO
	}
	type want struct {
		purchaseOrd models.PurchaseOrders
		err         error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When create a valid purchase order",
			setup: func() {
				_, err := factories.NewBuyerFactory(db).Create(models.Buyer{ID: 1})
				require.NoError(t, err)
				_, err = factories.NewProductRecordsFactory(db).Create(models.ProductRecord{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.PurchaseOrdersDTO{
					OrderNumber:     fixture.OrderNumber,
					OrderDate:       fixture.OrderDate,
					TrackingCode:    fixture.TrackingCode,
					BuyerID:         fixture.BuyerID,
					ProductRecordID: fixture.ProductRecordID,
				},
			},
			want: want{
				purchaseOrd: fixture,
				err:         nil,
			},
		},
		{
			name: "Error - When create a duplicated purchase order",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.PurchaseOrdersDTO{
					OrderNumber:     fixture.OrderNumber,
					OrderDate:       fixture.OrderDate,
					TrackingCode:    fixture.TrackingCode,
					BuyerID:         fixture.BuyerID,
					ProductRecordID: fixture.ProductRecordID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			rp := NewPurchaseOrdersRepository(db)

			got, err := rp.Create(tt.dto)
			log.Println(got, err)

			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want.purchaseOrd, got)
			}
		})
	}
}
