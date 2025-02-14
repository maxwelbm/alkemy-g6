package productsrp

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewProductFactory(db)
	prod := factory.Build(models.Product{ID: 1})
	newCode := "NEWCODE"

	type arg struct {
		id  int
		dto models.ProductDTO
	}
	type want struct {
		product models.Product
		err     error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully updating a Product",
			setup: func() {
				_, err := factory.Create(prod)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
				dto: models.ProductDTO{
					ProductCode: &newCode,
				},
			},
			want: want{
				product: func() models.Product {
					cpy := prod
					cpy.ProductCode = newCode
					return cpy
				}(),
				err: nil,
			},
		},
		{
			name: "Error - When the product is not found",
			arg: arg{
				id: 1,
				dto: models.ProductDTO{
					ProductCode: &newCode,
				},
			},
			want: want{
				err: models.ErrProductNotFound,
			},
		},
		{
			name: "Error - When trying to update a Product's SellerID to a non-existent Seller",
			setup: func() {
				_, err := factory.Create(prod)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
				dto: func() models.ProductDTO {
					undefinedSellerID := 10
					return models.ProductDTO{
						SellerID: &undefinedSellerID,
					}
				}(),
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			},
		},
		{
			name: "Error - When attempting to update a Product with a duplicate ProductCode",
			setup: func() {
				_, err := factory.Create(prod)
				require.NoError(t, err)
				_, err = factory.Create(factory.Build(models.Product{ID: 2}))
				require.NoError(t, err)
			},
			arg: arg{
				id: 2,
				dto: models.ProductDTO{
					ProductCode: &prod.ProductCode,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
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
			got, err := rp.Update(tt.id, tt.dto)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.product, got)
		})
	}
}
