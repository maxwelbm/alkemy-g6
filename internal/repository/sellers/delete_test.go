package sellersrp

import (
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
		id int
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
				id: fixture.ID,
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
				id: fixture.ID + 1,
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

			err := rp.Delete(tt.id)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
		})
	}
}
