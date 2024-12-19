package service

import (
	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
	"errors"
)

var (
	ErrSellerServiceProductsAssociated = errors.New("Cannot delete seller: products are still associated. Please remove or reassign associated products before deleting.")
)

type SellerDefault struct {
	repo repository.RepoDB
}

func NewSellerService(repositorySeller repository.RepoDB) *SellerDefault {
	return &SellerDefault{
		repo: repositorySeller,
	}
}

func (s *SellerDefault) GetAll() ([]modelsSeller.Seller, error) {
	return s.repo.SellersDB.GetAll()
}

func (s *SellerDefault) GetById(id int) (sel modelsSeller.Seller, err error) {
	return s.repo.SellersDB.GetById(id)
}

func (s *SellerDefault) GetByCid(cid int) (sel modelsSeller.Seller, err error) {
	return s.repo.SellersDB.GetByCid(cid)
}

func (s *SellerDefault) PostSeller(seller modelsSeller.Seller) error {
	return s.repo.SellersDB.PostSeller(seller)
}

func (s *SellerDefault) PatchSeller(seller modelsSeller.Seller) error {
	return s.repo.SellersDB.PatchSeller(seller)
}

func (s *SellerDefault) Delete(id int) (err error) {
	allProducts, err := s.repo.ProductsDB.GetAll()
	if err != nil {
		return
	}
	for _, product := range allProducts {
		if product.SellerID == id {
			err = ErrSellerServiceProductsAssociated
			return
		}
	}
	return s.repo.SellersDB.Delete(id)
}
