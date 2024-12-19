package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
)

type ProductsDefault struct {
	repo models.ProductRepository
}

func NewProductsDefault(repo models.ProductRepository) *ProductsDefault {
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

func (s *ProductsDefault) Update(id int, prod models.ProductDTO) (updatedProd models.Product, err error) {
	updatedProd, err = s.repo.Update(id, prod)
	return
}

func (s *ProductsDefault) Delete(id int) (err error) {
	err = s.repo.Delete(id)
	return
}
