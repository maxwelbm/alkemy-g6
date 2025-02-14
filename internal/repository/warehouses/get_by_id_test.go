package warehousesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewWarehouseFactory(db)
	wh := factory.Build(models.Warehouse{ID: 1})

	type arg struct {
		id int
	}
	type want struct {
		warehouses models.Warehouse
		err        error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When the warehouse is found",
			setup: func() {
				_, err := factory.Create(wh)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				warehouses: wh,
			},
		},
		{
			name: "When the warehouse is not found",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrWareHouseNotFound,
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
			got, err := rp.GetByID(tt.id)

			// Assert
			require.ErrorIs(t, err, tt.want.err)
			require.Equal(t, tt.want.warehouses, got)
		})
	}
}
