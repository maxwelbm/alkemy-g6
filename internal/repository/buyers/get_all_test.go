package buyersrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	db, tuncrate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewBuyerFactory(db)
	buyer1 := factory.Build(models.Buyer{ID: 1})
	buyer2 := factory.Build(models.Buyer{ID: 2})
	buyer3 := factory.Build(models.Buyer{ID: 3})

	type want struct {
		buyers []models.Buyer
	}
	tests := []struct {
		name  string
		setup func()
		want
	}{
		{
			name: "When buyers are registered",
			setup: func() {
				_, err := factory.Create(buyer1)
				require.NoError(t, err)
				_, err = factory.Create(buyer2)
				require.NoError(t, err)
				_, err = factory.Create(buyer3)
				require.NoError(t, err)
			},
			want: want{
				buyers: []models.Buyer{buyer1, buyer2, buyer3},
			},
		},
		{
			name: "When buyers are not registered",
			want: want{
				buyers: []models.Buyer(nil),
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

			got, err := rp.GetAll()

			require.NoError(t, err)
			require.Equal(t, tt.want.buyers, got)
		})
	}
}
