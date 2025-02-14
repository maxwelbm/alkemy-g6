package sellersrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetById(t *testing.T) {
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
			name: "When find Seller by id",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)

			},
			arg: arg{
				dto: models.SellerDTO{
					ID:          fixture.ID,
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
			name: "When not found Seller by id",
			setup: func() {
				_, err := factory.Create(fixture)
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
				err: errors.New("seller not found"),
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

			seller, err := rp.GetByID(tt.dto.ID)

			// Assert
			if tt.err != nil {
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want.seller, seller)
		})
	}
}
