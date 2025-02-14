package buyersrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	db, tuncrate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewBuyerFactory(db)
	fixture := factory.Build(models.Buyer{ID: 1})

	type arg struct {
		id int
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
			name: "When the buyer is found",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				buyer: fixture,
			},
		},
		{
			name: "When the buyer is not found",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrBuyerNotFound,
			},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(tuncrate)
			if tt.setup != nil {
				tt.setup()
			}

			rp := NewBuyersRepository(db)

			got, err := rp.GetByID(tt.id)

			require.ErrorIs(t, tt.want.err, err)
			require.Equal(t, tt.want.buyer, got)
		})
	}
}
