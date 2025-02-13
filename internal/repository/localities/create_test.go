package localitiesrp

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewLocalityFactory(db)
	fixture := factory.Build(models.Locality{ID: 1})

	type arg struct {
		dto models.LocalityDTO
	}
	type want struct {
		locality models.Locality
		err      error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully creating a new Locality",
			arg: arg{
				dto: models.LocalityDTO{
					LocalityName: &fixture.LocalityName,
					ProvinceName: &fixture.ProvinceName,
					CountryName:  &fixture.CountryName,
				},
			},
			want: want{
				locality: fixture,
				err:      nil,
			},
		},
		{
			name: "Error - When creating a duplicated Locality",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.LocalityDTO{
					LocalityName: &fixture.LocalityName,
					ProvinceName: &fixture.ProvinceName,
					CountryName:  &fixture.CountryName,
				},
			},
			want: want{
				err: errors.New("Duplicate entry"),
			},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			// Arrange
			rp := NewLocalityRepository(db)
			// Act
			got, err := rp.Create(tt.dto)

			// Assert
			if tt.err != nil {
				require.Contains(t, err.Error(), tt.err.Error())
			}
			require.Equal(t, tt.want.locality, got)
		})
	}
}
