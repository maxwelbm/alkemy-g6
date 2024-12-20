package service

import (
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository"
)

type ProductsDefault struct {
	repo repository.RepoDB
}

func NewProductsDefault(repo repository.RepoDB) *ProductsDefault {
	return &ProductsDefault{repo: repo}
}

func (s *ProductsDefault) GetAll() (list []models.Product, err error) {
	list, err = s.repo.ProductsDB.GetAll()
	return
}

func (s *ProductsDefault) GetById(id int) (prod models.Product, err error) {
	prod, err = s.repo.ProductsDB.GetById(id)
	return
}

func (s *ProductsDefault) Create(prod models.ProductDTO) (newProd models.Product, err error) {
	// Check if the seller exists
	if prod.SellerID != 0 {
		if _, err = s.repo.SellersDB.GetById(prod.SellerID); err != nil {
			err = errors.Join(errors.New("Seller not found"), err)
			return
		}
	}

	newProd, err = s.repo.ProductsDB.Create(prod)
	return
}

func (s *ProductsDefault) Update(id int, prod models.ProductDTO) (updatedProd models.Product, err error) {
	// Check if the seller exists
	if prod.SellerID != 0 {
		if _, err = s.repo.SellersDB.GetById(prod.SellerID); err != nil {
			err = errors.Join(errors.New("Seller not found"), err)
			return
		}
	}

	updatedProd, err = s.repo.ProductsDB.Update(id, prod)
	return
}

func (s *ProductsDefault) Delete(id int) (err error) {
	err = s.repo.ProductsDB.Delete(id)
	return
}
