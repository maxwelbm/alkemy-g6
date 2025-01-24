package service_test

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	productbatchesrp "github.com/maxwelbm/alkemy-g6/internal/repository/product_batches"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

func TestProductBatchesDefault_Create(t *testing.T) {
	prodBatchesFixture := models.ProductBatches{
		ID:                 1,
		BatchNumber:        "PB01",
		InitialQuantity:    100,
		CurrentQuantity:    50,
		CurrentTemperature: 20.0,
		MinimumTemperature: -5.0,
		DueDate:            "2022-04-04",
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  "08:00:00",
		ProductID:          1,
		SectionID:          1,
	}
	tests := []struct {
		name              string
		prodBatches       models.ProductBatches
		err               error
		wantedProdBatches models.ProductBatches
		wantedErr         error
	}{
		{
			name:              "When the repository returns a product batch",
			prodBatches:       prodBatchesFixture,
			err:               nil,
			wantedProdBatches: prodBatchesFixture,
			wantedErr:         nil,
		},
		{
			name:              "When the repository returns an error",
			prodBatches:       prodBatchesFixture,
			err:               &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedProdBatches: prodBatchesFixture,
			wantedErr:         &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := productbatchesrp.NewProductBatchesRepositoryMock()
			sv := service.NewProductBatchesService(rp)
			dto := models.ProductBatchesDTO{
				BatchNumber:        tt.wantedProdBatches.BatchNumber,
				InitialQuantity:    tt.prodBatches.InitialQuantity,
				CurrentQuantity:    tt.prodBatches.CurrentQuantity,
				CurrentTemperature: tt.prodBatches.CurrentTemperature,
				MinimumTemperature: tt.prodBatches.MinimumTemperature,
				DueDate:            tt.prodBatches.DueDate,
				ManufacturingDate:  tt.prodBatches.ManufacturingDate,
				ManufacturingHour:  tt.prodBatches.ManufacturingHour,
				ProductID:          tt.prodBatches.ProductID,
				SectionID:          tt.prodBatches.SectionID,
			}
			// Act
			rp.On("Create", dto).Return(tt.prodBatches, tt.err)
			prodBatches, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedProdBatches, prodBatches)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
