package models

type ProductRepository interface {
	GetAll() (list []Product, err error)
}
