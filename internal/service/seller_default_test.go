package service_test

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	sellersrp "github.com/maxwelbm/alkemy-g6/internal/repository/sellers"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSellersDefault_GetAll(t *testing.T) {
	rp := sellersrp.NewSellersRepositoryMock()
	rp.On("GetAll").Return([]models.Seller{
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
			ID:          1,
			CID:         "789",
			CompanyName: "Company C",
			Address:     "789 Street",
			Telephone:   "234567890",
			LocalityID:  3,
		},
	}, nil)

	s := service.NewSellersService(rp)
	sellers, err := s.GetAll()

	assert.NoError(t, err)
	assert.Len(t, sellers, 3)
	assert.Equal(t, "Company A", sellers[0].CompanyName)
	assert.Equal(t, "Company B", sellers[1].CompanyName)
	assert.Equal(t, "Company C", sellers[2].CompanyName)
}

func TestSellersDefault_GetByID(t *testing.T) {
	rp := sellersrp.NewSellersRepositoryMock()

	rp.On("GetByID", mock.AnythingOfType("int")).Return(models.Seller{
		ID:          1,
		CID:         "123",
		CompanyName: "Company A",
		Address:     "123 Street",
		Telephone:   "012345678",
		LocalityID:  1,
	}, nil)

	s := service.NewSellersService(rp)
	seller, err := s.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Company A", seller.CompanyName)
}

func TestSellersDefault_GetByCid(t *testing.T) {
	rp := sellersrp.NewSellersRepositoryMock()

	rp.On("GetByCid", mock.AnythingOfType("int")).Return(models.Seller{
		ID:          1,
		CID:         "123",
		CompanyName: "Company A",
		Address:     "123 Street",
		Telephone:   "012345678",
		LocalityID:  1,
	}, nil)

	s := service.NewSellersService(rp)
	seller, err := s.GetByCid(1)

	assert.NoError(t, err)
	assert.Equal(t, "123", seller.CID)
}

/*
func TestSellersDefault_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := models.NewMockSellersRepository(ctrl)
	mockRepo.EXPECT().Create(models.SellerDTO{}).Return(models.Seller{}, nil)

	s := service.NewSellersService(mockRepo)
	seller, err := s.Create(models.SellerDTO{})

	assert.NoError(t, err)
	assert.NotNil(t, seller)
}

func TestSellersDefault_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := models.NewMockSellersRepository(ctrl)
	mockRepo.EXPECT().Update(1, models.SellerDTO{}).Return(models.Seller{}, nil)

	s := service.NewSellersService(mockRepo)
	seller, err := s.Update(1, models.SellerDTO{})

	assert.NoError(t, err)
	assert.NotNil(t, seller)
}

func TestSellersDefault_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := models.NewMockSellersRepository(ctrl)
	mockRepo.EXPECT().Delete(1).Return(nil)

	s := service.NewSellersService(mockRepo)
	err := s.Delete(1)

	assert.NoError(t, err)
}

func TestSellersDefault_GetByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := models.NewMockSellersRepository(ctrl)
	mockRepo.EXPECT().GetByID(1).Return(models.Seller{}, errors.New("not found"))

	s := service.NewSellersService(mockRepo)
	seller, err := s.GetByID(1)

	assert.Error(t, err)
	assert.Equal(t, models.Seller{}, seller)
}
*/
