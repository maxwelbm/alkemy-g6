package models

type ProductService interface {
	GetAll() (list []Product, err error)
}
