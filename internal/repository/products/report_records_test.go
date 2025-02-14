package productsrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestReportRecords(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewProductFactory(db)
	prod1 := factory.Build(models.Product{ID: 1})
	prod2 := factory.Build(models.Product{ID: 2})

	type arg struct {
		id int
	}
	type want struct {
		reports []models.ProductReportRecords
		err     error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When reporting all products",
			setup: func() {
				_, err := factory.Create(prod1)
				require.NoError(t, err)
				_, err = factory.Create(prod2)
				require.NoError(t, err)
			},
			want: want{
				reports: []models.ProductReportRecords{
					{
						ProductID:    1,
						Description:  prod1.Description,
						RecordsCount: 0,
					},
					{
						ProductID:    2,
						Description:  prod2.Description,
						RecordsCount: 0,
					},
				},
			},
		},
		{
			name: "When reporting by product id",
			setup: func() {
				_, err := factory.Create(prod1)
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				reports: []models.ProductReportRecords{
					{
						ProductID:    1,
						Description:  prod1.Description,
						RecordsCount: 0,
					},
				},
			},
		},
		{
			name: "When reporting by product id with product records",
			setup: func() {
				_, err := factory.Create(prod1)
				require.NoError(t, err)
				_, err = factories.NewProductRecordsFactory(db).Create(models.ProductRecord{ProductID: prod1.ID})
				require.NoError(t, err)
			},
			arg: arg{
				id: 1,
			},
			want: want{
				reports: []models.ProductReportRecords{
					{
						ProductID:    1,
						Description:  prod1.Description,
						RecordsCount: 1,
					},
				},
			},
		},
		{
			name: "When reporting by an id that is Not Found",
			arg: arg{
				id: 1,
			},
			want: want{
				err: models.ErrReportRecordNotFound,
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
			rp := NewProducts(db)
			// Act
			got, err := rp.ReportRecords(tt.id)

			// Assert
			if tt.err != nil {
				require.Equal(t, tt.err.Error(), err.Error())
			}
			require.Equal(t, tt.want.reports, got)
		})
	}
}
