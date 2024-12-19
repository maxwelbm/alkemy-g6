package service

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

func (s *SellerDefault) GetByCid(cid int) (sel modelsSeller.Seller, err error) {
	return s.rp.GetByCid(cid)
}

func (s *SellerDefault) PostSeller(seller modelsSeller.Seller) error {
	return s.rp.PostSeller(seller)
}

func (s *SellerDefault) PatchSeller(seller modelsSeller.Seller) error {
	return s.rp.PatchSeller(seller)
}

func (s *SellerDefault) Delete(id int) (err error) {
	return s.rp.Delete(id)
}
