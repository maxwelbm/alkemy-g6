package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type ProductService struct {
	repo models.ProductRepository
}

func NewProductsService(repo models.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() (list []models.Product, err error) {
	list, err = s.repo.GetAll()
	return
}

func (s *ProductService) GetById(id int) (prod models.Product, err error) {
	prod, err = s.repo.GetById(id)
	return
}

func (s *ProductService) Create(prod models.ProductDTO) (newProd models.Product, err error) {
	newProd, err = s.repo.Create(prod)
	return
}

func (s *ProductService) Update(id int, prod models.ProductDTO) (updatedProd models.Product, err error) {
	updatedProd, err = s.repo.Update(id, prod)
	return
}

func (s *ProductService) Delete(id int) (err error) {
	err = s.repo.Delete(id)
	return
}
