package service_test

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	sectionsrp "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

var sectionsFixture = []models.Section{
	{
		ID:                 1,
		SectionNumber:      "Section 1",
		CurrentTemperature: 22.5,
		MinimumTemperature: 18.0,
		CurrentCapacity:    50,
		MinimumCapacity:    20,
		MaximumCapacity:    100,
		WarehouseID:        1,
		ProductTypeID:      1,
	},
	{
		ID:                 2,
		SectionNumber:      "Section 2",
		CurrentTemperature: 23.0,
		MinimumTemperature: 18.5,
		CurrentCapacity:    60,
		MinimumCapacity:    25,
		MaximumCapacity:    110,
		WarehouseID:        2,
		ProductTypeID:      2,
	},
	{
		ID:                 3,
		SectionNumber:      "Section 3",
		CurrentTemperature: 21.0,
		MinimumTemperature: 17.0,
		CurrentCapacity:    70,
		MinimumCapacity:    30,
		MaximumCapacity:    120,
		WarehouseID:        3,
		ProductTypeID:      3,
	},
}

func TestSectionsDefault_GetAll(t *testing.T) {
	tests := []struct {
		name          string
		section       []models.Section
		err           error
		wantedSection []models.Section
		wantedErr     error
	}{
		{
			name:          "When the repository returns a section",
			section:       sectionsFixture,
			err:           nil,
			wantedSection: sectionsFixture,
			wantedErr:     nil,
		},
		{
			name:          "When the repository returns an error",
			section:       []models.Section{},
			err:           errors.New("internal error"),
			wantedSection: []models.Section{},
			wantedErr:     errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			rp.On("GetAll").Return(tt.section, tt.err)
			sv := service.NewSectionsService(rp)

			// Act
			section, err := sv.GetAll()

			// Assert
			require.Equal(t, tt.wantedSection, section)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSectionsDefault_GetByID(t *testing.T) {
	tests := []struct {
		name          string
		section       models.Section
		err           error
		wantedSection models.Section
		wantedErr     error
	}{
		{
			name:          "When the repository returns a section",
			section:       sectionsFixture[0],
			err:           nil,
			wantedSection: sectionsFixture[0],
			wantedErr:     nil,
		},
		{
			name:          "When the repository returns an error",
			section:       models.Section{},
			err:           models.ErrSectionNotFound,
			wantedSection: models.Section{},
			wantedErr:     models.ErrSectionNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			rp.On("GetByID", tt.section.ID).Return(tt.section, tt.err)
			sv := service.NewSectionsService(rp)

			// Act
			section, err := sv.GetByID(tt.section.ID)

			// Assert
			require.Equal(t, tt.wantedSection, section)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSectionsDefault_Create(t *testing.T) {
	tests := []struct {
		name          string
		section       models.Section
		err           error
		wantedSection models.Section
		wantedErr     error
	}{
		{
			name:          "When the repository returns a section",
			section:       sectionsFixture[0],
			err:           nil,
			wantedSection: sectionsFixture[0],
			wantedErr:     nil,
		},
		{
			name:          "When the repository returns an error",
			section:       sectionsFixture[0],
			err:           &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedSection: sectionsFixture[0],
			wantedErr:     &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			sv := service.NewSectionsService(rp)
			dto := models.SectionDTO{
				SectionNumber:      &tt.section.SectionNumber,
				CurrentTemperature: &tt.section.CurrentTemperature,
				MinimumTemperature: &tt.section.MinimumTemperature,
				CurrentCapacity:    &tt.section.CurrentCapacity,
				MinimumCapacity:    &tt.section.MinimumCapacity,
				MaximumCapacity:    &tt.section.MaximumCapacity,
				WarehouseID:        &tt.section.WarehouseID,
				ProductTypeID:      &tt.section.ProductTypeID,
			}
			// Act
			rp.On("Create", dto).Return(tt.section, tt.err)
			section, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedSection, section)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSectionsDefault_Update(t *testing.T) {
	tests := []struct {
		name          string
		section       models.Section
		err           error
		wantedSection models.Section
		wantedErr     error
	}{
		{
			name:          "When the repository returns a section",
			section:       sectionsFixture[0],
			err:           nil,
			wantedSection: sectionsFixture[0],
			wantedErr:     nil,
		},
		{
			name:          "When the repository returns an error",
			section:       sectionsFixture[0],
			err:           models.ErrSectionNotFound,
			wantedSection: sectionsFixture[0],
			wantedErr:     models.ErrSectionNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			sv := service.NewSectionsService(rp)
			dto := models.SectionDTO{
				SectionNumber:      &tt.section.SectionNumber,
				CurrentTemperature: &tt.section.CurrentTemperature,
				MinimumTemperature: &tt.section.MinimumTemperature,
				CurrentCapacity:    &tt.section.CurrentCapacity,
				MinimumCapacity:    &tt.section.MinimumCapacity,
				MaximumCapacity:    &tt.section.MaximumCapacity,
				WarehouseID:        &tt.section.WarehouseID,
				ProductTypeID:      &tt.section.ProductTypeID,
			}
			// Act
			rp.On("Update", 1, dto).Return(tt.section, tt.err)
			section, err := sv.Update(1, dto)

			// Assert
			require.Equal(t, tt.wantedSection, section)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSectionsDefault_Delete(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		wantedErr error
	}{
		{
			name:      "When the repository deletes the section sucessfully",
			err:       nil,
			wantedErr: nil,
		},
		{
			name:      "When the repository returns an error",
			err:       models.ErrSectionNotFound,
			wantedErr: models.ErrSectionNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			sv := service.NewSectionsService(rp)

			// Act
			rp.On("Delete", 1).Return(tt.err)
			err := sv.Delete(1)

			// Assert
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestSectionsDefault_ReportProducts(t *testing.T) {
	tests := []struct {
		name          string
		reports       []models.ProductReport
		err           error
		wantedReports []models.ProductReport
		wantedErr     error
	}{
		{
			name:          "When the repository returns the report",
			reports:       []models.ProductReport{{SectionID: 1, ProductsCount: 1, SectionNumber: "A"}},
			err:           nil,
			wantedReports: []models.ProductReport{{SectionID: 1, ProductsCount: 1, SectionNumber: "A"}},
			wantedErr:     nil,
		},
		{
			name:          "When the repository returns an error",
			reports:       []models.ProductReport{},
			err:           errors.New("internal error"),
			wantedReports: []models.ProductReport{},
			wantedErr:     errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := sectionsrp.NewSectionsRepositoryMock()
			sv := service.NewSectionsService(rp)

			// Act
			rp.On("ReportProducts", 0).Return(tt.reports, tt.err)
			pr, err := sv.ReportProducts(0)

			// Assert
			require.Equal(t, tt.wantedReports, pr)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
