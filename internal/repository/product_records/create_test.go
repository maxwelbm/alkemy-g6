package productrecordsrp

import (
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
	factory := factories.NewProductRecordsFactory(db)
	fixture := factory.Build(models.ProductRecord{ID: 1})

	type arg struct {
		dto models.ProductRecordDTO
	}
	type want struct {
		productRecord models.ProductRecord
		err           error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully creating a new product record",
			setup: func() {
				_, err := factories.NewProductFactory(db).Create(models.Product{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.ProductRecordDTO{
					LastUpdateDate: &fixture.LastUpdateDate,
					PurchasePrice:  &fixture.PurchasePrice,
					SalePrice:      &fixture.SalePrice,
					ProductID:      &fixture.ProductID,
				},
			},
			want: want{
				productRecord: fixture,
				err:           nil,
			},
		},
		{
			name: "Error - When creating a productRecord from a non-existent product",
			arg: arg{
				dto: models.ProductRecordDTO{
					LastUpdateDate: &fixture.LastUpdateDate,
					PurchasePrice:  &fixture.PurchasePrice,
					SalePrice:      &fixture.SalePrice,
					ProductID:      &fixture.ProductID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
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
			rp := NewProductRecordsRepository(db)
			// Act
			got, err := rp.Create(tt.dto)
			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.productRecord, got)
		})
	}
}
