package sectionsrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	factory := factories.NewSectionFactory(db)
	section1 := factory.Build(models.Section{ID: 1})
	section2 := factory.Build(models.Section{ID: 2})
	section3 := factory.Build(models.Section{ID: 3})

	type want struct {
		sections []models.Section
	}
	tests := []struct {
		name  string
		setup func()
		want
	}{
		{
			name: "When sections are registered",
			setup: func() {
				_, err := factory.Create(section1)
				require.NoError(t, err)
				_, err = factory.Create(section2)
				require.NoError(t, err)
				_, err = factory.Create(section3)
				require.NoError(t, err)
			},
			want: want{
				sections: []models.Section{section1, section2, section3},
			},
		},
		{
			name: "When no sections are registered",
			want: want{
				sections: []models.Section(nil),
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
			got, err := rp.GetAll()

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.want.sections, got)
		})
	}
}
