package warehousesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewWarehouseFactory(db)
	wh := factory.Build(models.Warehouse{ID: 1})
	newCode := "NEWCODE"

	type arg struct {
		id  int
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
			name: "When successfully updating a warehouse",
			setup: func() {
				_, err := factory.Create(wh)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
				dto: models.WarehouseDTO{
					WarehouseCode: &newCode,
				},
			},
			want: want{
				warehouse: func() models.Warehouse {
					cpy := wh
					cpy.WarehouseCode = newCode
					return cpy
				}(),
				err: nil,
			},
		},
		{
			name: "Error - When the warehouse is not found",
			arg: arg{
				id: 1,
				dto: models.WarehouseDTO{
					WarehouseCode: &newCode,
				},
			},
			want: want{
				err: models.ErrWareHouseNotFound,
			},
		},
		{
			name: "Error - When attempting to update a Warehouse with a duplicate WarehouseCode",
			setup: func() {
				_, err := factory.Create(wh)
				require.NoError(t, err)
				_, err = factory.Create(factory.Build(models.Warehouse{ID: 2}))
				require.NoError(t, err)
			},
			arg: arg{
				id: 2,
				dto: models.WarehouseDTO{
					WarehouseCode: &wh.WarehouseCode,
				},
			},
			want: want{
				err: models.ErrWareHouseCodeExist,
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
			got, err := rp.Update(tt.id, tt.dto)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.warehouse, got)
		})
	}
}
