package sectionsrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewSectionFactory(db)
	section := factory.Build(models.Section{ID: 1})

	type arg struct {
		id int
	}
	type want struct {
		warehouses models.Section
		err        error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When the warehouse is found",
			setup: func() {
				_, err := factory.Create(section)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				warehouses: section,
			},
		},
		{
			name: "When the warehouse is not found",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrSectionNotFound,
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
			rp := NewSectionsRepository(db)
			// Act
			got, err := rp.GetByID(tt.id)

			// Assert
			require.ErrorIs(t, err, tt.want.err)
			require.Equal(t, tt.want.warehouses, got)
		})
	}
}
