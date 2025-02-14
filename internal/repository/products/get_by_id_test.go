package productsrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewProductFactory(db)
	prod := factory.Build(models.Product{ID: 1})

	type arg struct {
		id int
	}
	type want struct {
		products models.Product
		err      error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When the product is found",
			setup: func() {
				_, err := factory.Create(prod)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				products: prod,
			},
		},
		{
			name: "When the product is not found",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrProductNotFound,
			},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			// Arrange
			rp := NewProducts(db)
			// Act
			got, err := rp.GetByID(tt.id)

			// Assert
			require.ErrorIs(t, err, tt.want.err)
			require.Equal(t, tt.want.products, got)
		})
	}
}
