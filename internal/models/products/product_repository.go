package models

type ProductRepository interface {
	GetAll() (list []Product, err error)
	GetById(id int) (prod Product, err error)
	Create(prod ProductDTO) (newProd Product, err error)
	Update(id int, prod ProductDTO) (updatedProd Product, err error)
	Delete(id int) (err error)
}
