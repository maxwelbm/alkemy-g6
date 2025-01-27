package service_test

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	buyersrp "github.com/maxwelbm/alkemy-g6/internal/repository/buyers"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var buyersFixture = []models.Buyer{
	{
		ID:           1,
		CardNumberID: "123456789",
		FirstName:    "Cleiton",
		LastName:     "Ortega",
	},
	{
		ID:           1,
		CardNumberID: "012345678",
		FirstName:    "Joao",
		LastName:     "Pedro",
	},
}

var buyerPurchaseOrdersReportFixture = []models.BuyerPurchaseOrdersReport{
	{
		ID:                  1,
		CardNumberID:        "123456789",
		FirstName:           "Petra",
		LastName:            "Grunheidt",
		PurchaseOrdersCount: 100,
	},
	{
		ID:                  2,
		CardNumberID:        "012345678",
		FirstName:           "Vitoria",
		LastName:            "Vital",
		PurchaseOrdersCount: 101,
	},
	{
		ID:                  3,
		CardNumberID:        "901234567",
		FirstName:           "Izabelly",
		LastName:            "Melo",
		PurchaseOrdersCount: 102,
	},
}

func TestBuyersDefault_GetAll(t *testing.T) {
	tests := []struct {
		name        string
		buyer       []models.Buyer
		err         error
		wantedBuyer []models.Buyer
		wantedErr   error
	}{
		{
			name:        "When the repository returns a buyer",
			buyer:       buyersFixture,
			err:         nil,
			wantedBuyer: buyersFixture,
			wantedErr:   nil,
		},
		{
			name:        "When the repository returns an error",
			buyer:       []models.Buyer{},
			err:         errors.New("internal error"),
			wantedBuyer: []models.Buyer{},
			wantedErr:   errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := buyersrp.NewBuyersRepositoryMock()
			rp.On("GetAll").Return(tt.buyer, tt.err)
			sv := service.NewBuyersService(rp)

			// Act
			buyer, err := sv.GetAll()

			// Assert
			require.Equal(t, tt.wantedBuyer, buyer)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestBuyersDefault_GetByID(t *testing.T) {
	tests := []struct {
		name        string
		buyer       models.Buyer
		err         error
		wantedBuyer models.Buyer
		wantedErr   error
	}{
		{
			name:        "When the repository returns a buyer",
			buyer:       buyersFixture[0],
			err:         nil,
			wantedBuyer: buyersFixture[0],
			wantedErr:   nil,
		},
		{
			name:        "When the repository returns an error",
			buyer:       models.Buyer{},
			err:         models.ErrSellerNotFound,
			wantedBuyer: models.Buyer{},
			wantedErr:   models.ErrSellerNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := buyersrp.NewBuyersRepositoryMock()
			rp.On("GetByID", tt.buyer.ID).Return(tt.buyer, tt.err)
			sv := service.NewBuyersService(rp)

			// Act
			buyer, err := sv.GetByID(tt.buyer.ID)

			// Assert
			require.Equal(t, tt.wantedBuyer, buyer)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestBuyersDefault_GetByCardNumberID(t *testing.T) {
	tests := []struct {
		name        string
		buyer       models.Buyer
		err         error
		wantedBuyer models.Buyer
		wantedErr   error
	}{
		{
			name:        "When the repository returns a buyer by cid",
			buyer:       buyersFixture[0],
			err:         nil,
			wantedBuyer: buyersFixture[0],
			wantedErr:   nil,
		},
		{
			name:        "When the repository returns an error",
			buyer:       models.Buyer{},
			err:         models.ErrSellerNotFound,
			wantedBuyer: models.Buyer{},
			wantedErr:   models.ErrSellerNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := buyersrp.NewBuyersRepositoryMock()
			rp.On("GetByCardNumberID", tt.buyer.CardNumberID).Return(tt.buyer, tt.err)
			sv := service.NewBuyersService(rp)

			// Act
			buyer, err := sv.GetByCardNumberID(tt.buyer.CardNumberID)

			// Assert
			require.Equal(t, tt.wantedBuyer, buyer)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestBuyersDefault_Create(t *testing.T) {
	tests := []struct {
		name        string
		buyer       models.Buyer
		err         error
		wantedBuyer models.Buyer
		wantedErr   error
	}{
		{
			name:        "When the repository returns a buyer",
			buyer:       buyersFixture[0],
			err:         nil,
			wantedBuyer: buyersFixture[0],
			wantedErr:   nil,
		},
		{
			name:        "When the repository returns an error",
			buyer:       buyersFixture[0],
			err:         &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedBuyer: buyersFixture[0],
			wantedErr:   &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := buyersrp.NewBuyersRepositoryMock()
			sv := service.NewBuyersService(rp)
			dto := models.BuyerDTO{
				ID: &tt.buyer.ID,
			}
			// Act
			rp.On("Create", dto).Return(tt.buyer, tt.err)
			buyer, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedBuyer, buyer)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestBuyersDefault_Update(t *testing.T) {
	tests := []struct {
		name         string
		buyer        models.Buyer
		err          error
		wantedSeller models.Buyer
		wantedErr    error
	}{
		{
			name:         "When the repository returns a buyer",
			buyer:        buyersFixture[0],
			err:          nil,
			wantedSeller: buyersFixture[0],
			wantedErr:    nil,
		},
		{
			name:         "When the repository returns an error",
			buyer:        buyersFixture[0],
			err:          &mysql.MySQLError{Number: mysqlerr.CodeCannotDeleteOrUpdateParentRow},
			wantedSeller: buyersFixture[0],
			wantedErr:    &mysql.MySQLError{Number: mysqlerr.CodeCannotDeleteOrUpdateParentRow},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := buyersrp.NewBuyersRepositoryMock()
			sv := service.NewBuyersService(rp)
			dto := models.BuyerDTO{
				ID:           &tt.buyer.ID,
				CardNumberID: &tt.buyer.CardNumberID,
				FirstName:    &tt.buyer.FirstName,
				LastName:     &tt.buyer.LastName,
			}
			// Act
			rp.On("Update", 1, dto).Return(tt.buyer, tt.err)
			buyer, err := sv.Update(*dto.ID, dto)

			// Assert
			require.Equal(t, tt.wantedSeller, buyer)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestBuyersDefault_Delete(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		wantedErr error
	}{
		{
			name:      "When the repository deletes the buyer sucessfully",
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
			rp := buyersrp.NewBuyersRepositoryMock()
			sv := service.NewBuyersService(rp)

			// Act
			rp.On("Delete", 1).Return(tt.err)
			err := sv.Delete(1)

			// Assert
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestBuyersDefault_ReportPurchaseOrders(t *testing.T) {
	tests := []struct {
		name                            string
		buyerPurchaseOrdersReport       []models.BuyerPurchaseOrdersReport
		err                             error
		wantedBuyerPurchaseOrdersReport []models.BuyerPurchaseOrdersReport
		wantedErr                       error
	}{
		{
			name:                            "When the repository returns a buyer",
			buyerPurchaseOrdersReport:       buyerPurchaseOrdersReportFixture,
			err:                             nil,
			wantedBuyerPurchaseOrdersReport: buyerPurchaseOrdersReportFixture,
			wantedErr:                       nil,
		},
		{
			name:                            "When the repository returns an error",
			buyerPurchaseOrdersReport:       []models.BuyerPurchaseOrdersReport{},
			err:                             models.ErrBuyerNotFound,
			wantedBuyerPurchaseOrdersReport: []models.BuyerPurchaseOrdersReport{},
			wantedErr:                       models.ErrBuyerNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := buyersrp.NewBuyersRepositoryMock()
			rp.On("ReportPurchaseOrders", mock.AnythingOfType("int")).Return(tt.buyerPurchaseOrdersReport, tt.err)
			sv := service.NewBuyersService(rp)

			// Act
			reportPurchaseOrdersList, err := sv.ReportPurchaseOrders(1)

			// Assert
			require.Equal(t, tt.wantedBuyerPurchaseOrdersReport, reportPurchaseOrdersList)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
