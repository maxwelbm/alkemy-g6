package service_test

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	productsrp "github.com/maxwelbm/alkemy-g6/internal/repository/products"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

var productsFixture = []models.Product{
	{
		ID:             1,
		ProductCode:    "P001",
		Description:    "Product 1",
		Height:         10,
		Length:         20,
		Width:          30,
		NetWeight:      40,
		ExpirationRate: 1,
		FreezingRate:   2,
		RecomFreezTemp: -10,
		ProductTypeID:  1,
		SellerID:       1,
	},
	{
		ID:             2,
		ProductCode:    "P002",
		Description:    "Product 2",
		Height:         11,
		Length:         21,
		Width:          31,
		NetWeight:      41,
		ExpirationRate: 2,
		FreezingRate:   3,
		RecomFreezTemp: -11,
		ProductTypeID:  2,
		SellerID:       2,
	},
	{
		ID:             3,
		ProductCode:    "P003",
		Description:    "Product 3",
		Height:         12,
		Length:         22,
		Width:          32,
		NetWeight:      42,
		ExpirationRate: 3,
		FreezingRate:   4,
		RecomFreezTemp: -12,
		ProductTypeID:  3,
		SellerID:       3,
	},
}

func TestProductsDefault_GetAll(t *testing.T) {

	tests := []struct {
		name            string
		product         []models.Product
		err             error
		expectedProduct []models.Product
		expectedErr     error
	}{
		{
			name:            "When the repository returns a product",
			product:         productsFixture,
			err:             nil,
			expectedProduct: productsFixture,
			expectedErr:     nil,
		},
		{
			name:            "When the repository returns an error",
			product:         []models.Product{},
			err:             errors.New("internal error"),
			expectedProduct: []models.Product{},
			expectedErr:     errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := productsrp.NewProductsRepositoryMock()
			rp.On("GetAll").Return(tt.product, tt.err)
			sv := service.NewProductsService(rp)

			// Act
			result, err := sv.GetAll()

			// Assert
			require.Equal(t, tt.expectedProduct, result)
			require.Equal(t, tt.expectedErr, err)
		})
	}

}

func TestProductsDefault_GetByID(t *testing.T) {
	tests := []struct {
		name         string
		product      models.Product
		err          error
		expecProduct models.Product
		expectedErr  error
	}{
		{
			name:         "When the repository returns a product",
			product:      productsFixture[0],
			err:          nil,
			expecProduct: productsFixture[0],
			expectedErr:  nil,
		},
		{
			name:         "When the repository returns an error",
			product:      models.Product{},
			err:          models.ErrProductNotFound,
			expecProduct: models.Product{},
			expectedErr:  models.ErrProductNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := productsrp.NewProductsRepositoryMock()
			rp.On("GetByID", tt.product.ID).Return(tt.product, tt.err)
			sv := service.NewProductsService(rp)

			// Act
			result, err := sv.GetByID(tt.product.ID)

			// Assert
			require.Equal(t, tt.expecProduct, result)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestProductsDefault_ReportRecords(t *testing.T) {
	tests := []struct {
		name          string
		productRecord []models.ProductReportRecords
		err           error
		expecProduct  []models.ProductReportRecords
		expectedErr   error
	}{
		{
			name: "When the repository returns a product record",
			productRecord: []models.ProductReportRecords{
				{
					ProductID:    1,
					Description:  "Product 1",
					RecordsCount: 1,
				},
			},
			err: nil,
			expecProduct: []models.ProductReportRecords{
				{
					ProductID:    1,
					Description:  "Product 1",
					RecordsCount: 1,
				},
			},
			expectedErr: nil,
		},
		{
			name:          "When the repository returns an error",
			productRecord: []models.ProductReportRecords{},
			err:           errors.New("internal error"),
			expecProduct:  []models.ProductReportRecords{},
			expectedErr:   errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := productsrp.NewProductsRepositoryMock()
			sv := service.NewProductsService(rp)

			// Act
			rp.On("ReportRecords", 0).Return(tt.productRecord, tt.err)
			result, err := sv.ReportRecords(0)

			// Assert
			require.Equal(t, tt.expecProduct, result)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestProductsDefault_Create(t *testing.T) {
	tests := []struct {
		name         string
		product      models.Product
		err          error
		expecProduct models.Product
		expectedErr  error
	}{
		{
			name:         "When the repository returns a product",
			product:      productsFixture[0],
			err:          nil,
			expecProduct: productsFixture[0],
			expectedErr:  nil,
		},
		{
			name:         "When the repository returns an error",
			product:      productsFixture[0],
			err:          &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			expecProduct: productsFixture[0],
			expectedErr:  &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := productsrp.NewProductsRepositoryMock()
			sv := service.NewProductsService(rp)
			dto := models.ProductDTO{
				ProductCode:    tt.product.ProductCode,
				Description:    tt.product.Description,
				Height:         tt.product.Height,
				Length:         tt.product.Length,
				Width:          tt.product.Width,
				NetWeight:      tt.product.NetWeight,
				ExpirationRate: tt.product.ExpirationRate,
				FreezingRate:   tt.product.FreezingRate,
				RecomFreezTemp: tt.product.RecomFreezTemp,
				ProductTypeID:  tt.product.ProductTypeID,
				SellerID:       tt.product.SellerID,
			}

			// Act
			rp.On("Create", dto).Return(tt.product, tt.err)
			result, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.expecProduct, result)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestProductsDefault_Update(t *testing.T) {
	tests := []struct {
		name         string
		product      models.Product
		err          error
		expecProduct models.Product
		expectedErr  error
	}{
		{
			name:         "When the repository returns a product",
			product:      productsFixture[0],
			err:          nil,
			expecProduct: productsFixture[0],
			expectedErr:  nil,
		},
		{
			name:         "When the repository returns an error",
			product:      productsFixture[0],
			err:          models.ErrProductNotFound,
			expecProduct: productsFixture[0],
			expectedErr:  models.ErrProductNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := productsrp.NewProductsRepositoryMock()
			sv := service.NewProductsService(rp)
			dto := models.ProductDTO{
				ID:             tt.product.ID,
				ProductCode:    tt.product.ProductCode,
				Description:    tt.product.Description,
				Height:         tt.product.Height,
				Length:         tt.product.Length,
				Width:          tt.product.Width,
				NetWeight:      tt.product.NetWeight,
				ExpirationRate: tt.product.ExpirationRate,
				FreezingRate:   tt.product.FreezingRate,
				RecomFreezTemp: tt.product.RecomFreezTemp,
				ProductTypeID:  tt.product.ProductTypeID,
				SellerID:       tt.product.SellerID,
			}

			// Act
			rp.On("Update", tt.product.ID, dto).Return(tt.product, tt.err)
			result, err := sv.Update(tt.product.ID, dto)

			// Assert
			require.Equal(t, tt.expecProduct, result)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestProductsDefault_Delete(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		expectedErr error
	}{
		{
			name:        "When the repository returns a product",
			err:         nil,
			expectedErr: nil,
		},
		{
			name:        "When the repository returns an error",
			err:         models.ErrProductNotFound,
			expectedErr: models.ErrProductNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := productsrp.NewProductsRepositoryMock()
			sv := service.NewProductsService(rp)

			// Act
			rp.On("Delete", 1).Return(tt.err)
			err := sv.Delete(1)

			// Assert
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
