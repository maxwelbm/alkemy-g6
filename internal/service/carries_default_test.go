package service_test

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	carriesrp "github.com/maxwelbm/alkemy-g6/internal/repository/carries"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

var carryFixture = models.Carry{
	ID:          1,
	CID:         "123",
	CompanyName: "Company A",
	Address:     "123 Street",
	PhoneNumber: "012345678",
	LocalityID:  1,
}

func TestCarriesDefault_Create(t *testing.T) {
	tests := []struct {
		name        string
		carry       models.Carry
		err         error
		wantedCarry models.Carry
		wantedErr   error
	}{
		{
			name:        "When the repository returns a carry",
			carry:       carryFixture,
			err:         nil,
			wantedCarry: carryFixture,
			wantedErr:   nil,
		},
		{
			name:        "When the repository returns an error",
			carry:       carryFixture,
			err:         &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedCarry: carryFixture,
			wantedErr:   &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := carriesrp.NewCarriesRepositoryMock()
			sv := service.NewCarriesService(rp)
			dto := models.CarryDTO{
				CID:         &tt.carry.CID,
				CompanyName: &tt.carry.CompanyName,
				Address:     &tt.carry.Address,
				PhoneNumber: &tt.carry.PhoneNumber,
				LocalityID:  &tt.carry.LocalityID,
			}
			// Act
			rp.On("Create", dto).Return(tt.carry, tt.err)
			carry, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedCarry, carry)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
