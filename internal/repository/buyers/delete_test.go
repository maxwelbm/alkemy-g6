package buyersrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	db, tuncrate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewBuyerFactory(db)

	type arg struct {
		id int
	}
	type want struct {
		err error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully deleting a Buyer",
			setup: func() {
				_, err := factory.Create(models.Buyer{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "Error - When trying to delete a buyer that does not exist",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrBuyerNotFound,
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

			err := rp.Delete(tt.id)

			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
		})
	}
}
