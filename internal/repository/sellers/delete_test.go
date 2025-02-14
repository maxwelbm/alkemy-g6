package sellersrp

import (
	"fmt"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
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
			name: "When delete a Seller",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.SellerDTO{
					ID:          fixture.ID,
					CompanyName: fixture.CompanyName,
					Address:     fixture.Address,
					Telephone:   fixture.Telephone,
					LocalityID:  fixture.LocalityID,
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "Error - When Seller not found to delete",
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
				err: models.ErrSellerNotFound,
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

			fmt.Printf("tt.name: %v\n", tt.name)
			fmt.Printf("tt.dto.ID: %v\n", tt.dto.ID)
			err := rp.Delete(tt.dto.ID)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
		})
	}
}
