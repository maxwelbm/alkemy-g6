package buyersrp

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db, tuncrate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewBuyerFactory(db)
	fixture := factory.Build(models.Buyer{ID: 1})

	type arg struct {
		dto models.BuyerDTO
	}
	type want struct {
		buyer models.Buyer
		err   error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully creating a new Buyer",
			arg: arg{
				dto: models.BuyerDTO{
					CardNumberID: &fixture.CardNumberID,
					FirstName:    &fixture.FirstName,
					LastName:     &fixture.LastName,
				},
			},
			want: want{
				buyer: fixture,
				err:   nil,
			},
		},
		{
			name: "Error - When creating a duplicated Buyer",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.BuyerDTO{
					CardNumberID: &fixture.CardNumberID,
					FirstName:    &fixture.FirstName,
					LastName:     &fixture.LastName,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(tuncrate)
			if tt.setup != nil {
				tt.setup()
			}

			rp := NewBuyersRepository(db)

			got, err := rp.Create(tt.dto)

			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.buyer, got)
		})
	}
}
