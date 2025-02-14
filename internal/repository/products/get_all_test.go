package productsrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewProductFactory(db)
	prod1 := factory.Build(models.Product{ID: 1})
	prod2 := factory.Build(models.Product{ID: 2})

	type want struct {
		products []models.Product
	}
	tests := []struct {
		name  string
		setup func()
		want
	}{
		{
			name: "When products are registered",
			setup: func() {
				_, err := factory.Create(prod1)
				require.NoError(t, err)
				_, err = factory.Create(prod2)
				require.NoError(t, err)
			},
			want: want{
				products: []models.Product{prod1, prod2},
			},
		},
		{
			name: "When no products are registered",
			want: want{
				products: []models.Product(nil),
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
			got, err := rp.GetAll()

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.want.products, got)
		})
	}
}
