package sectionsrp

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewSectionFactory(db)
	fixture := factory.Build(models.Section{ID: 1})
	type arg struct {
		dto models.SectionDTO
	}
	type want struct {
		section models.Section
		err     error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully creating a new section",
			setup: func() {
				_, err := factories.NewWarehouseFactory(db).Create(models.Warehouse{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.SectionDTO{
					SectionNumber:      &fixture.SectionNumber,
					CurrentTemperature: &fixture.CurrentTemperature,
					MinimumTemperature: &fixture.MinimumTemperature,
					CurrentCapacity:    &fixture.CurrentCapacity,
					MinimumCapacity:    &fixture.MinimumCapacity,
					MaximumCapacity:    &fixture.MaximumCapacity,
					WarehouseID:        &fixture.WarehouseID,
					ProductTypeID:      &fixture.ProductTypeID,
				},
			},
			want: want{
				section: fixture,
				err:     nil,
			},
		},
		{
			name: "Error - When creating a section from a non-existent warehouse",
			arg: arg{
				dto: models.SectionDTO{
					SectionNumber:      &fixture.SectionNumber,
					CurrentTemperature: &fixture.CurrentTemperature,
					MinimumTemperature: &fixture.MinimumTemperature,
					CurrentCapacity:    &fixture.CurrentCapacity,
					MinimumCapacity:    &fixture.MinimumCapacity,
					MaximumCapacity:    &fixture.MaximumCapacity,
					WarehouseID:        &fixture.WarehouseID,
					ProductTypeID:      &fixture.ProductTypeID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			},
		},
		{
			name: "Error - When creating a duplicated section",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.SectionDTO{
					SectionNumber:      &fixture.SectionNumber,
					CurrentTemperature: &fixture.CurrentTemperature,
					MinimumTemperature: &fixture.MinimumTemperature,
					CurrentCapacity:    &fixture.CurrentCapacity,
					MinimumCapacity:    &fixture.MinimumCapacity,
					MaximumCapacity:    &fixture.MaximumCapacity,
					WarehouseID:        &fixture.WarehouseID,
					ProductTypeID:      &fixture.ProductTypeID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
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
			got, err := rp.Create(tt.dto)
			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.section, got)
		})
	}
}
	