package productbatchesrp

import (
	"strings"
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
	factory := factories.NewProductBatchesFactory(db)
	fixture := factory.Build(models.ProductBatches{ID: 1})

	type arg struct {
		dto models.ProductBatchesDTO
	}
	type want struct {
		productBatches models.ProductBatches
		err            error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When create a new Product Batches",
			setup: func() {
				_, err := factories.NewProductFactory(db).Create(models.Product{ID: 1})
				require.NoError(t, err)

				_, err = factories.NewSectionFactory(db).Create(models.Section{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.ProductBatchesDTO{
					BatchNumber:        fixture.BatchNumber,
					InitialQuantity:    fixture.InitialQuantity,
					CurrentQuantity:    fixture.CurrentQuantity,
					CurrentTemperature: fixture.CurrentTemperature,
					MinimumTemperature: fixture.MinimumTemperature,
					DueDate:            fixture.DueDate,
					ManufacturingDate:  fixture.ManufacturingDate,
					ManufacturingHour:  fixture.ManufacturingHour,
					ProductID:          fixture.ProductID,
					SectionID:          fixture.SectionID,
				},
			},
			want: want{
				productBatches: fixture,
				err:            nil,
			},
		},
		{
			name: "Error - When try create a new Product Batches and dependences not exists",
			arg: arg{
				dto: models.ProductBatchesDTO{
					BatchNumber:        fixture.BatchNumber,
					InitialQuantity:    fixture.InitialQuantity,
					CurrentQuantity:    fixture.CurrentQuantity,
					CurrentTemperature: fixture.CurrentTemperature,
					MinimumTemperature: fixture.MinimumTemperature,
					DueDate:            fixture.DueDate,
					ManufacturingDate:  fixture.ManufacturingDate,
					ManufacturingHour:  fixture.ManufacturingHour,
					ProductID:          fixture.ProductID,
					SectionID:          fixture.SectionID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
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
			rp := NewProductBatchesRepository(db)
			// Act
			got, err := rp.Create(tt.dto)

			// Format dates
			got.DueDate = formatDateToUse(got.DueDate)
			got.ManufacturingDate = formatDateToUse(got.ManufacturingDate)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.productBatches, got)
		})
	}
}

func formatDateToUse(date string) string {
	if date == "" {
		return ""
	}
	return strings.Split(date, "T")[0]
}
