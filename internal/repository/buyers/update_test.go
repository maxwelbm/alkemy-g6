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

func TestUpdate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewBuyerFactory(db)
	fixture := factory.Build(models.Buyer{ID: 1})
	newCardNumber := "NEWCODE"

	type arg struct {
		id  int
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
			name: "When successfully updating a Buyer",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
				dto: models.BuyerDTO{
					CardNumberID: &newCardNumber,
				},
			},
			want: want{
				buyer: func() models.Buyer {
					cpy := fixture
					cpy.CardNumberID = newCardNumber
					return cpy
				}(),
				err: nil,
			},
		},
		{
			name: "Error - When the buyer is not found",
			arg: arg{
				id: 1,
				dto: models.BuyerDTO{
					CardNumberID: &newCardNumber,
				},
			},
			want: want{
				err: models.ErrBuyerNotFound,
			},
		},
		{
			name: "Error - When attempting to update a Buyer with a duplicate BuyerCode",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
				_, err = factory.Create(factory.Build(models.Buyer{ID: 2}))
				require.NoError(t, err)
			},
			arg: arg{
				id: 2,
				dto: models.BuyerDTO{
					CardNumberID: &fixture.CardNumberID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			rp := NewBuyersRepository(db)

			got, err := rp.Update(tt.id, tt.dto)

			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.buyer, got)
		})
	}
}
