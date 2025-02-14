package sellersrp

import (
	"fmt"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewSellerFactory(db)
	fixture := factory.Build(models.Seller{ID: 1})

	type arg struct {
		dto models.SellerDTO
	}
	type want struct {
		seller models.Seller
		err    error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When create a new Seller",
			setup: func() {
				err := factory.CheckLocalityExists(fixture.LocalityID)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.SellerDTO{
					ID:          fixture.ID + 1,
					CID:         fixture.CID,
					CompanyName: fixture.CompanyName,
					Address:     fixture.Address,
					Telephone:   fixture.Telephone,
					LocalityID:  fixture.LocalityID,
				},
			},
			want: want{
				seller: fixture,
				err:    nil,
			},
		},
		{
			name: "Error - When try create new empty Seller",
			arg: arg{
				dto: models.SellerDTO{
					CID:         fixture.CID,
					CompanyName: fixture.CompanyName,
					Address:     fixture.Address,
					Telephone:   fixture.Telephone,
					LocalityID:  fixture.LocalityID,
				},
			},
			want: want{
				err: errors.New("Error 1452"),
			},
		},
		{
			name: "Error - When creating a duplicated Seller",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.SellerDTO{
					CID:         fixture.CID,
					CompanyName: fixture.CompanyName,
					Address:     fixture.Address,
					Telephone:   fixture.Telephone,
					LocalityID:  fixture.LocalityID,
				},
			},
			want: want{
				err: errors.New("Error 1062"),
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
			rp := NewSellersRepository(db)
			// Act

			fmt.Printf("tt.dto: %v\n", tt.dto)
			got, err := rp.Create(tt.dto)

			fmt.Printf("got: %v\n", got)
			fmt.Printf("tt.want.seller: %v\n", tt.want.seller)
			// Assert
			if tt.err != nil {
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want.seller, got)
		})
	}
}
