package buyersrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestReportPurchaseOrders(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewBuyerFactory(db)
	buyer1 := factory.Build(models.Buyer{ID: 1})
	buyer2 := factory.Build(models.Buyer{ID: 2})

	type arg struct {
		id int
	}
	type want struct {
		reports []models.BuyerPurchaseOrdersReport
		err     error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When reporting all buyers",
			setup: func() {
				_, err := factory.Create(buyer1)
				require.NoError(t, err)
				_, err = factory.Create(buyer2)
				require.NoError(t, err)
			},
			want: want{
				reports: []models.BuyerPurchaseOrdersReport{
					{
						ID:                  1,
						CardNumberID:        buyer1.CardNumberID,
						FirstName:           buyer1.FirstName,
						LastName:            buyer1.LastName,
						PurchaseOrdersCount: 0,
					},
					{
						ID:                  2,
						CardNumberID:        buyer2.CardNumberID,
						FirstName:           buyer2.FirstName,
						LastName:            buyer2.LastName,
						PurchaseOrdersCount: 0,
					},
				},
			},
		},
		{
			name: "When reporting by product id",
			setup: func() {
				_, err := factory.Create(buyer1)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				reports: []models.BuyerPurchaseOrdersReport{
					{
						ID:                  1,
						CardNumberID:        buyer1.CardNumberID,
						FirstName:           buyer1.FirstName,
						LastName:            buyer1.LastName,
						PurchaseOrdersCount: 0,
					},
				},
			},
		},
		{
			name: "When reporting by an id that is Not Found",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrBuyerNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			rp := NewBuyersRepository(db)
			// Act
			got, err := rp.ReportPurchaseOrders(tt.id)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.reports, got)
		})
	}
}
