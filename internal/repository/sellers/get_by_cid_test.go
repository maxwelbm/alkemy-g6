package sellersrp

import (
	"strconv"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetByCid(t *testing.T) {
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
			name: "When find Seller by cid",
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
			name: "When not found Seller by cid",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)

			},
			arg: arg{
				dto: models.SellerDTO{
					ID:          fixture.ID,
					CID:         randstr.Alphanumeric(10),
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

			cidInt, _ := strconv.Atoi(tt.arg.dto.CID)
			seller, err := rp.GetByCid(cidInt)

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
