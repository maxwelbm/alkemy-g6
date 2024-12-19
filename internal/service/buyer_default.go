package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
)

type BuyerDefault struct {
	repo repository.RepoDB
}

func NewBuyerService(repo repository.RepoDB) *BuyerDefault {
	return &BuyerDefault{
		repo: repo,
	}
}

func (s *BuyerDefault) GetAll() ([]models.Buyer, error) {
	return s.repo.BuyersDB.GetAll()
}

func (s *BuyerDefault) GetById(id int) (models.Buyer, error) {
	return s.repo.BuyersDB.GetById(id)
}

func (s *BuyerDefault) GetByCardNumberId(cardNumberId string) (models.Buyer, error) {
	return s.repo.BuyersDB.GetByCardNumberId(cardNumberId)
}

func (s *BuyerDefault) PostBuyer(buyer models.Buyer) (models.Buyer, error) {
	return s.repo.BuyersDB.PostBuyer(buyer)
}

func (s *BuyerDefault) PatchBuyer(buyer models.BuyerDTO) (models.Buyer, error) {
	return s.repo.BuyersDB.PatchBuyer(buyer)
}

func (s *BuyerDefault) Delete(id int) (err error) {
	return s.repo.BuyersDB.Delete(id)
}
