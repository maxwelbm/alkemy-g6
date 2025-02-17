package sectionsrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestReportProducts(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	sectionFactory := factories.NewSectionFactory(db)
	productBatchFactory := factories.NewProductBatchesFactory(db)

	section1 := sectionFactory.Build(models.Section{ID: 1, SectionNumber: "S001"})
	section2 := sectionFactory.Build(models.Section{ID: 2, SectionNumber: "S002"})

	type arg struct {
		sectionID int
	}

	type want struct {
		reports []models.ProductReport
		err     error
	}

	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When getting report for specific section with products",
			setup: func() {
				_, err := sectionFactory.Create(section1)
				require.NoError(t, err)

				for i := 0; i < 3; i++ {
					pb := models.ProductBatches{
						SectionID: section1.ID,
					}
					_, err := productBatchFactory.Create(pb)
					require.NoError(t, err)
				}
			},
			arg: arg{
				sectionID: 1,
			},
			want: want{
				reports: []models.ProductReport{
					{
						SectionID:     1,
						SectionNumber: "S001",
						ProductsCount: 3,
					},
				},
			},
		},
		{
			name: "When getting report for all sections",
			setup: func() {
				// Create sections
				_, err := sectionFactory.Create(section1)
				require.NoError(t, err)
				_, err = sectionFactory.Create(section2)
				require.NoError(t, err)

				// Create product batches for section1
				for i := 0; i < 3; i++ {
					pb := models.ProductBatches{
						SectionID: section1.ID,
					}
					_, err := productBatchFactory.Create(pb)
					require.NoError(t, err)
				}

				// Create product batches for section2
				for i := 0; i < 2; i++ {
					pb := models.ProductBatches{
						SectionID: section2.ID,
					}
					_, err := productBatchFactory.Create(pb)
					require.NoError(t, err)
				}
			},
			arg: arg{
				sectionID: 0,
			},
			want: want{
				reports: []models.ProductReport{
					{
						SectionID:     1,
						SectionNumber: "S001",
						ProductsCount: 3,
					},
					{
						SectionID:     2,
						SectionNumber: "S002",
						ProductsCount: 2,
					},
				},
			},
		},
		{
			name: "When section does not exist",
			arg: arg{
				sectionID: 999,
			},
			want: want{
				err:     models.ErrSectionNotFound,
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
			got, err := rp.ReportProducts(tt.arg.sectionID)

			// Assert
			if tt.want.err != nil {
				require.ErrorIs(t, err, tt.want.err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want.reports, got)
		})
	}
}
