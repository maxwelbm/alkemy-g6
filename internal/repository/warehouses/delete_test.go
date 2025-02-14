package warehousesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewWarehouseFactory(db)

	type arg struct {
		id int
	}
	type want struct {
		err error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully deleting a new warehouse",
			setup: func() {
				_, err := factory.Create(models.Warehouse{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "Error - When trying to delete a warehouse that does not exist",
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
			err := rp.Delete(tt.id)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
		})
	}
}
