package models

type ProductService interface {
	GetAll() (list []Product, err error)
	GetById(id int) (prod Product, err error)
	Delete(id int) (err error)
	Create(prod ProductDTO) (newProd Product, err error)
}
