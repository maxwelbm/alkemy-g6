package productsrp

import (
	"log"
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
	factory := factories.NewProductFactory(db)
	fixture := factory.Build(models.Product{ID: 1})

	type arg struct {
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
			name: "When successfully creating a new Product",
			setup: func() {
				_, err := factories.NewSellerFactory(db).Create(models.Seller{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.ProductDTO{
					ProductCode:    &fixture.ProductCode,
					Description:    &fixture.Description,
					Height:         &fixture.Height,
					Length:         &fixture.Length,
					Width:          &fixture.Width,
					NetWeight:      &fixture.NetWeight,
					ExpirationRate: &fixture.ExpirationRate,
					FreezingRate:   &fixture.FreezingRate,
					RecomFreezTemp: &fixture.RecomFreezTemp,
					ProductTypeID:  &fixture.ProductTypeID,
					SellerID:       &fixture.SellerID,
				},
			},
			want: want{
				product: fixture,
				err:     nil,
			},
		},
		{
			name: "Error - When creating a Product from a non-existent Seller",
			arg: arg{
				dto: models.ProductDTO{
					ProductCode:    &fixture.ProductCode,
					Description:    &fixture.Description,
					Height:         &fixture.Height,
					Length:         &fixture.Length,
					Width:          &fixture.Width,
					NetWeight:      &fixture.NetWeight,
					ExpirationRate: &fixture.ExpirationRate,
					FreezingRate:   &fixture.FreezingRate,
					RecomFreezTemp: &fixture.RecomFreezTemp,
					ProductTypeID:  &fixture.ProductTypeID,
					SellerID:       &fixture.SellerID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			},
		},
		{
			name: "Error - When creating a duplicated Product",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.ProductDTO{
					ProductCode:    &fixture.ProductCode,
					Description:    &fixture.Description,
					Height:         &fixture.Height,
					Length:         &fixture.Length,
					Width:          &fixture.Width,
					NetWeight:      &fixture.NetWeight,
					ExpirationRate: &fixture.ExpirationRate,
					FreezingRate:   &fixture.FreezingRate,
					RecomFreezTemp: &fixture.RecomFreezTemp,
					ProductTypeID:  &fixture.ProductTypeID,
					SellerID:       &fixture.SellerID,
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
			log.Println(*tt.dto.Description)
			got, err := rp.Create(tt.dto)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.product, got)
		})
	}
}
