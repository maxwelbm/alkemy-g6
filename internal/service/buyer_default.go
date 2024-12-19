package service

import (
	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

type BuyerDefault struct {
	rp modelsBuyer.BuyerRepository
}

func NewBuyerService(repositoryBuyer modelsBuyer.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{
		rp: repositoryBuyer,
	}
}

func (s *BuyerDefault) GetAll() ([]modelsBuyer.Buyer, error) {
	return s.rp.GetAll()
}

func (s *BuyerDefault) GetById(id int) (modelsBuyer.Buyer, error) {
	return s.rp.GetById(id)
}

func (s *BuyerDefault) GetByCardNumberId(cardNumberId string) (modelsBuyer.Buyer, error) {
	return s.rp.GetByCardNumberId(cardNumberId)
}

func (s *BuyerDefault) PostBuyer(buyer modelsBuyer.Buyer) (modelsBuyer.Buyer, error) {
	return s.rp.PostBuyer(buyer)
}

func (s *BuyerDefault) PatchBuyer(buyer modelsBuyer.BuyerDTO) (modelsBuyer.Buyer, error) {
	return s.rp.PatchBuyer(buyer)
}

func (s *BuyerDefault) Delete(id int) (err error) {
	return s.rp.Delete(id)
}
