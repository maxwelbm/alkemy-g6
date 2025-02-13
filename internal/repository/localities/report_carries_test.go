package localitiesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestReportCarries(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	type arg struct {
		id int
	}
	type want struct {
		locality []models.LocalityCarriesReport
		err      error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When reporting all localities",
			setup: func() {
				factory := factories.NewLocalityFactory(db)
				_, err := factory.Create(models.Locality{
					LocalityName: "Sao Paulo",
					ProvinceName: "Sao Paulo",
					CountryName:  "Brazil",
				})
				require.NoError(t, err)
				_, err = factory.Create(models.Locality{
					LocalityName: "Rio de Janeiro",
					ProvinceName: "Rio de Janeiro",
					CountryName:  "Brazil",
				})
				require.NoError(t, err)
			},
			want: want{
				locality: []models.LocalityCarriesReport{
					{
						ID:           1,
						LocalityName: "Sao Paulo",
						CarriesCount: 0,
					},
					{
						ID:           2,
						LocalityName: "Rio de Janeiro",
						CarriesCount: 0,
					},
				},
				err: nil,
			},
		},
		{
			name: "When reporting one locality by id",
			setup: func() {
				factory := factories.NewLocalityFactory(db)
				_, err := factory.Create(models.Locality{
					LocalityName: "Sao Paulo",
					ProvinceName: "Sao Paulo",
					CountryName:  "Brazil",
				})
				require.NoError(t, err)
				_, err = factory.Create(models.Locality{
					LocalityName: "Rio de Janeiro",
					ProvinceName: "Rio de Janeiro",
					CountryName:  "Brazil",
				})
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				locality: []models.LocalityCarriesReport{
					{
						ID:           1,
						LocalityName: "Sao Paulo",
						CarriesCount: 0,
					},
				},
				err: nil,
			},
		},
		{
			name: "When reporting one locality by id and the locality has sellers",
			setup: func() {
				factory := factories.NewLocalityFactory(db)
				_, err := factory.Create(models.Locality{
					LocalityName: "Sao Paulo",
					ProvinceName: "Sao Paulo",
					CountryName:  "Brazil",
				})
				require.NoError(t, err)
				sellerFactory := factories.NewCarryFactory(db)
				_, err = sellerFactory.Create(models.Carry{LocalityID: 1})
				require.NoError(t, err)
				_, err = sellerFactory.Create(models.Carry{LocalityID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				locality: []models.LocalityCarriesReport{
					{
						ID:           1,
						LocalityName: "Sao Paulo",
						CarriesCount: 2,
					},
				},
				err: nil,
			},
		},
		{
			name: "When reporting by id and locality is Not Found",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrLocalityNotFound,
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
			got, err := rp.ReportCarries(tt.id)

			// Assert
			if tt.err != nil {
				require.Equal(t, tt.err.Error(), err.Error())
			}
			require.Equal(t, tt.want.locality, got)
		})
	}
}
