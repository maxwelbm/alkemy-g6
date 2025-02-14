package carriesrp

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	models "github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewCarryFactory(db)
	carry := factory.Build(models.Carry{ID: 1})

	type arg struct {
		dto models.CarryDTO
	}
	type want struct {
		product models.Carry
		err     error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When successfully creating a new Carry",
			setup: func() {
				_, err := factories.NewLocalityFactory(db).Create(models.Locality{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.CarryDTO{
					CID:         &carry.CID,
					CompanyName: &carry.CompanyName,
					Address:     &carry.Address,
					PhoneNumber: &carry.PhoneNumber,
					LocalityID:  &carry.LocalityID,
				},
			},
			want: want{
				product: carry,
				err:     nil,
			},
		},
		{
			name: "Error - When creating a Carry from a non-existent Locality",
			arg: arg{
				dto: models.CarryDTO{
					CID:         &carry.CID,
					CompanyName: &carry.CompanyName,
					Address:     &carry.Address,
					PhoneNumber: &carry.PhoneNumber,
					LocalityID:  &carry.LocalityID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
			},
		},
		{
			name: "Error - When creating a Carry with a duplicated CID",
			setup: func() {
				_, err := factory.Create(carry)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.CarryDTO{
					CID:         &carry.CID,
					CompanyName: &carry.CompanyName,
					Address:     &carry.Address,
					PhoneNumber: &carry.PhoneNumber,
					LocalityID:  &carry.LocalityID,
				},
			},
			want: want{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
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
			rp := NewCarriesRepository(db)
			// Act
			got, err := rp.Create(tt.dto)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			}
			require.Equal(t, tt.want.product, got)
		})
	}
}
