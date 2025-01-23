package service_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	sellersrp "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

var sellersFixture = []models.Seller{
	{
		ID:          1,
		CID:         "123",
		CompanyName: "Company A",
		Address:     "123 Street",
		Telephone:   "012345678",
		LocalityID:  1,
	}, {
		ID:          2,
		CID:         "456",
		CompanyName: "Company B",
		Address:     "456 Street",
		Telephone:   "123456789",
		LocalityID:  2,
	}, {
		ID:          3,
		CID:         "789",
		CompanyName: "Company C",
		Address:     "789 Street",
		Telephone:   "234567890",
		LocalityID:  3,
	},
}

func TestSellersDefault_GetAll(t *testing.T) {
	tests := []struct {
		name         string
		seller       []models.Seller
		err          error
		wantedSeller []models.Seller
		wantedErr    error
	}{
		{
			name:         "When the repository returns a seller",
			seller:       sellersFixture,
			err:          nil,
			wantedSeller: sellersFixture,
			wantedErr:    nil,
		},
		{
			name:         "When the repository returns an error",
			seller:       []models.Seller{},
			err:          errors.New("internal error"),
			wantedSeller: []models.Seller{},
			wantedErr:    errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sellersrp.NewSellersRepositoryMock()
			rp.On("GetAll").Return(tt.seller, tt.err)
			sv := service.NewSellersService(rp)

			// Act
			seller, err := sv.GetAll()

			// Assert
			require.Equal(t, tt.wantedSeller, seller)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSellersDefault_GetByID(t *testing.T) {
	tests := []struct {
		name         string
		seller       models.Seller
		err          error
		wantedSeller models.Seller
		wantedErr    error
	}{
		{
			name:         "When the repository returns a seller",
			seller:       sellersFixture[0],
			err:          nil,
			wantedSeller: sellersFixture[0],
			wantedErr:    nil,
		},
		{
			name:         "When the repository returns an error",
			seller:       models.Seller{},
			err:          models.ErrSellerNotFound,
			wantedSeller: models.Seller{},
			wantedErr:    models.ErrSellerNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sellersrp.NewSellersRepositoryMock()
			rp.On("GetByID", tt.seller.ID).Return(tt.seller, tt.err)
			sv := service.NewSellersService(rp)

			// Act
			seller, err := sv.GetByID(tt.seller.ID)

			// Assert
			require.Equal(t, tt.wantedSeller, seller)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSellersDefault_GetByCid(t *testing.T) {
	tests := []struct {
		name         string
		seller       models.Seller
		err          error
		wantedSeller models.Seller
		wantedErr    error
	}{
		{
			name:         "When the repository returns a seller by cid",
			seller:       sellersFixture[0],
			err:          nil,
			wantedSeller: sellersFixture[0],
			wantedErr:    nil,
		},
		{
			name:         "When the repository returns an error",
			seller:       models.Seller{},
			err:          models.ErrSellerNotFound,
			wantedSeller: models.Seller{},
			wantedErr:    models.ErrSellerNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cidInt, _ := strconv.Atoi(tt.seller.CID)
			// Arrange
			rp := sellersrp.NewSellersRepositoryMock()
			rp.On("GetByCid", cidInt).Return(tt.seller, tt.err)
			sv := service.NewSellersService(rp)

			// Act
			seller, err := sv.GetByCid(cidInt)

			// Assert
			require.Equal(t, tt.wantedSeller, seller)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSellersDefault_Create(t *testing.T) {
	tests := []struct {
		name         string
		seller       models.Seller
		err          error
		wantedSeller models.Seller
		wantedErr    error
	}{
		{
			name:         "When the repository returns a seller",
			seller:       sellersFixture[0],
			err:          nil,
			wantedSeller: sellersFixture[0],
			wantedErr:    nil,
		},
		{
			name:         "When the repository returns an error",
			seller:       sellersFixture[0],
			err:          &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedSeller: sellersFixture[0],
			wantedErr:    &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sellersrp.NewSellersRepositoryMock()
			sv := service.NewSellersService(rp)
			dto := models.SellerDTO{
				ID:          tt.seller.ID,
				CID:         tt.seller.CID,
				CompanyName: tt.seller.CompanyName,
				Address:     tt.seller.Address,
				Telephone:   tt.seller.Telephone,
				LocalityID:  tt.seller.LocalityID,
			}
			// Act
			rp.On("Create", dto).Return(tt.seller, tt.err)
			seller, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedSeller, seller)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
