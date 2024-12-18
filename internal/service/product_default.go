package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/products"
)

type ProductsDefault struct {
	repo repository.Products
}

func NewProductsDefault(repo repository.Products) *ProductsDefault {
	return &ProductsDefault{repo: repo}
}

func (s *ProductsDefault) GetAll() (list []models.Product, err error) {
	list, err = s.repo.GetAll()
	return
}

func (s *ProductsDefault) GetById(id int) (prod models.Product, err error) {
	prod, err = s.repo.GetById(id)
	return
}

func (s *ProductsDefault) Create(prod models.ProductDTO) (newProd models.Product, err error) {
	newProd, err = s.repo.Create(prod)
	return
}
