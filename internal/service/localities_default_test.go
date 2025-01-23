package service_test

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	localitiesrp "github.com/maxwelbm/alkemy-g6/internal/repository/localities"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var localityFixture = models.Locality{
	LocalityName: "São Paulo",
	ProvinceName: "São Paulo",
	CountryName:  "Brazil",
}

func TestLocalitiesDefault_Create(t *testing.T) {
	tests := []struct {
		name           string
		locality       models.Locality
		err            error
		wantedLocality models.Locality
		wantedErr      error
	}{
		{
			name:           "When the repository returns a locality",
			locality:       localityFixture,
			err:            nil,
			wantedLocality: localityFixture,
			wantedErr:      nil,
		},
		{
			name:           "When the repository returns an error",
			locality:       localityFixture,
			err:            &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedLocality: localityFixture,
			wantedErr:      &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := localitiesrp.NewLocalitiesRepositoryMock()
			sv := service.NewLocalitiesService(rp)
			dto := models.LocalityDTO{
				LocalityName: &tt.locality.LocalityName,
				ProvinceName: &tt.locality.ProvinceName,
				CountryName:  &tt.locality.CountryName,
			}
			// Act
			rp.On("Create", dto).Return(tt.locality, tt.err)
			locality, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedLocality, locality)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestLocalitiesDefault_ReportCarries(t *testing.T) {
	tests := []struct {
		name         string
		report       []models.LocalityCarriesReport
		err          error
		wantedReport []models.LocalityCarriesReport
		wantedErr    error
	}{
		{
			name:         "When the repository returns a report",
			report:       []models.LocalityCarriesReport{{ID: 1, LocalityName: "São Paulo", CarriesCount: 1}},
			err:          nil,
			wantedReport: []models.LocalityCarriesReport{{ID: 1, LocalityName: "São Paulo", CarriesCount: 1}},
			wantedErr:    nil,
		},
		{
			name:         "When the repository returns an error",
			report:       []models.LocalityCarriesReport{},
			err:          &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedReport: []models.LocalityCarriesReport{},
			wantedErr:    &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := localitiesrp.NewLocalitiesRepositoryMock()
			sv := service.NewLocalitiesService(rp)
			// Act
			rp.On("ReportCarries", mock.AnythingOfType("int")).Return(tt.report, tt.err)
			report, err := sv.ReportCarries(1)

			// Assert
			require.Equal(t, tt.wantedReport, report)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestLocalitiesDefault_ReportSellers(t *testing.T) {
	tests := []struct {
		name         string
		report       []models.LocalitySellersReport
		err          error
		wantedReport []models.LocalitySellersReport
		wantedErr    error
	}{
		{
			name:         "When the repository returns a report",
			report:       []models.LocalitySellersReport{{ID: 1, LocalityName: "São Paulo", SellersCount: 1}},
			err:          nil,
			wantedReport: []models.LocalitySellersReport{{ID: 1, LocalityName: "São Paulo", SellersCount: 1}},
			wantedErr:    nil,
		},
		{
			name:         "When the repository returns an error",
			report:       []models.LocalitySellersReport{},
			err:          &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedReport: []models.LocalitySellersReport{},
			wantedErr:    &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := localitiesrp.NewLocalitiesRepositoryMock()
			sv := service.NewLocalitiesService(rp)
			// Act
			rp.On("ReportSellers", mock.AnythingOfType("int")).Return(tt.report, tt.err)
			report, err := sv.ReportSellers(1)

			// Assert
			require.Equal(t, tt.wantedReport, report)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
