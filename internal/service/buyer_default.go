package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type BuyerDefault struct {
	rp models.BuyerRepository
}

func NewBuyerService(rp models.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{
		rp: rp,
	}
}

func (s *BuyerDefault) GetAll() (buyers []models.Buyer, err error) {
	buyers, err = s.rp.GetAll()
	return
}

func (s *BuyerDefault) GetById(id int) (buyer models.Buyer, err error) {
	buyer, err = s.rp.GetById(id)
	return
}

func (s *BuyerDefault) GetByCardNumberId(cardNumberId string) (buyer models.Buyer, err error) {
	buyer, err = s.rp.GetByCardNumberId(cardNumberId)
	return
}

func (s *BuyerDefault) Create(buyer models.BuyerDTO) (buyerReturn models.Buyer, err error) {
	buyerReturn, err = s.rp.Create(buyer)
	return
}

func (s *BuyerDefault) Update(buyer models.BuyerDTO) (buyerReturn models.Buyer, err error) {
	buyerReturn, err = s.rp.Update(buyer)
	return
}

func (s *BuyerDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}
