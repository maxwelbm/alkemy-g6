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
				dto: func() models.LocalityDTO {
					locality := "Sao Paulo"
					province := "Sao Paulo"
					country := "Brazil"

					return models.LocalityDTO{
						LocalityName: &locality,
						ProvinceName: &province,
						CountryName:  &country,
					}
				}(),
			},
			want: want{
				locality: models.Locality{
					ID:           1,
					LocalityName: "Sao Paulo",
					ProvinceName: "Sao Paulo",
					CountryName:  "Brazil",
				},
				err: nil,
			},
		},
		{
			name: "Error - When creating a duplicated Locality",
			setup: func() {
				locality := models.Locality{LocalityName: "Sao Paulo", ProvinceName: "Sao Paulo", CountryName: "Brazil"}
				_, err := factories.NewLocalityFactory(db).Create(locality)
				require.NoError(t, err)
			},
			arg: arg{
				dto: func() models.LocalityDTO {
					locality := "Sao Paulo"
					province := "Sao Paulo"
					country := "Brazil"

					return models.LocalityDTO{
						LocalityName: &locality,
						ProvinceName: &province,
						CountryName:  &country,
					}
				}(),
			},
			want: want{
				err: errors.New("Duplicate entry"),
			},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

			// Cleans up sql entries
			truncate()
		})
	}
}
