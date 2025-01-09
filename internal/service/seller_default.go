package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type SellersDefault struct {
	rp models.SellersRepository
}

func NewSellersService(rp models.SellersRepository) *SellersDefault {
	return &SellersDefault{
		rp: rp,
	}
}

func (s *SellersDefault) GetAll() ([]models.Seller, error) {
	return s.rp.GetAll()
}

func (s *SellersDefault) GetById(id int) (seller models.Seller, err error) {
	return s.rp.GetById(id)
}

func (s *SellersDefault) GetByCid(cid int) (seller models.Seller, err error) {
	return s.rp.GetByCid(cid)
}

func (s *SellersDefault) Create(seller models.SellerDTO) (sellerToReturn models.Seller, err error) {
	return s.rp.Create(seller)
}

func (s *SellersDefault) Update(id int, seller models.SellerDTO) (sellerToReturn models.Seller, err error) {
	return s.rp.Update(id, seller)
}

func (s *SellersDefault) Delete(id int) (err error) {
	return s.rp.Delete(id)
}
