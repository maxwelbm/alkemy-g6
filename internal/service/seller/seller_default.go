package sellerService

import (
	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
)

type SellerDefault struct {
	rp modelsSeller.SellerRepository
}

func NewSellerService(repositorySeller modelsSeller.SellerRepository) *SellerDefault {
	return &SellerDefault{
		rp: repositorySeller,
	}
}

func (s *SellerDefault) GetAll() ([]modelsSeller.Seller, error) {
	return s.rp.GetAll()
}

func (s *SellerDefault) GetById(id int) (sel modelsSeller.Seller, err error) {
	return s.rp.GetById(id)
}
