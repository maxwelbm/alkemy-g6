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

func (s *SellerDefault) GetAll() (sellers []modelsSeller.Seller, err error) {
	sellers, err = s.rp.GetAll()
	return
}
