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
