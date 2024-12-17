package sellerService

import (
	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
)

func NewSellerService(repositorySeller modelsSeller.SellerRepository) *SellerDefault {
	return &SellerDefault{
		rp: repositorySeller,
	}
}

type SellerDefault struct {
	rp modelsSeller.SellerRepository
}

func (s *SellerDefault) FindAll() (sellers map[int]modelsSeller.Seller, err error) {
	return s.rp.FindAll()
}
