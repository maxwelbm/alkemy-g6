package warehousesrp

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewWarehouseFactory(db)
	fixture := factory.Build(models.Warehouse{ID: 1})

	type arg struct {
		dto models.WarehouseDTO
	}
	type want struct {
		warehouse models.Warehouse
		err       error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully creating a new Warehouse",
			arg: arg{
				dto: models.WarehouseDTO{
					Address:            &fixture.Address,
					Telephone:          &fixture.Telephone,
					WarehouseCode:      &fixture.WarehouseCode,
					MinimumCapacity:    &fixture.MinimumCapacity,
					MinimumTemperature: &fixture.MinimumTemperature,
				},
			},
			want: want{
				warehouse: fixture,
				err:       nil,
			},
		},
		{
			name: "Error - When creating a duplicated Warehouse",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.WarehouseDTO{
					Address:            &fixture.Address,
					Telephone:          &fixture.Telephone,
					WarehouseCode:      &fixture.WarehouseCode,
					MinimumCapacity:    &fixture.MinimumCapacity,
					MinimumTemperature: &fixture.MinimumTemperature,
				},
			},
			want: want{
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
			rp := NewWarehouseRepository(db)
			// Act
			got, err := rp.Create(tt.dto)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.warehouse, got)
		})
	}
}
