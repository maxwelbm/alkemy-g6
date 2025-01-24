package service_test

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	productrecordsrp "github.com/maxwelbm/alkemy-g6/internal/repository/product_records"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/require"
)

func TestProductRecordsDefault_Create(t *testing.T) {
	newProductRecord := models.ProductRecord{
		ID:             4,
		LastUpdateDate: "2021-09-01",
		PurchasePrice:  10.5,
		SalePrice:      15.5,
		ProductID:      1,
	}

	tests := []struct {
		name                  string
		productRecord         models.ProductRecord
		err                   error
		expectedProductRecord models.ProductRecord
		expectedErr           error
	}{
		{
			name:                  "Successfully create a new productRecord",
			productRecord:         newProductRecord,
			err:                   nil,
			expectedProductRecord: newProductRecord,
			expectedErr:           nil,
		},
		{
			name:                  "Error when trying to create a new productRecord",
			productRecord:         models.ProductRecord{},
			err:                   errors.New("internal error"),
			expectedProductRecord: models.ProductRecord{},
			expectedErr:           errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRecordDTO := models.ProductRecordDTO{
				ID:             tt.productRecord.ID,
				LastUpdateDate: tt.productRecord.LastUpdateDate,
				PurchasePrice:  tt.productRecord.PurchasePrice,
				SalePrice:      tt.productRecord.SalePrice,
				ProductID:      tt.productRecord.ProductID,
			}

			rp := productrecordsrp.NewProductRecordsRepositoryMock()
			rp.On("Create", productRecordDTO).Return(tt.productRecord, tt.err)
			sv := service.NewProductRecordsService(rp)

			productRecord, err := sv.Create(productRecordDTO)

			require.Equal(t, tt.expectedProductRecord, productRecord)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
