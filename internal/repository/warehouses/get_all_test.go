package warehousesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	factory := factories.NewWarehouseFactory(db)
	wh1 := factory.Build(models.Warehouse{ID: 1})
	wh2 := factory.Build(models.Warehouse{ID: 2})
	wh3 := factory.Build(models.Warehouse{ID: 3})

	type want struct {
		warehouses []models.Warehouse
	}

	tests := []struct {
		name  string
		setup func()
		want
	}{
		{
			name: "When warehouses are registered",
			setup: func() {
				_, err := factory.Create(wh1)
				require.NoError(t, err)
				_, err = factory.Create(wh2)
				require.NoError(t, err)
				_, err = factory.Create(wh3)
				require.NoError(t, err)
			},
			want: want{
				warehouses: []models.Warehouse{wh1, wh2, wh3},
			},
		},
		{
			name: "When no warehouses are registered",
			want: want{
				warehouses: []models.Warehouse(nil),
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
			got, err := rp.GetAll()

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.want.warehouses, got)
		})
	}
}
