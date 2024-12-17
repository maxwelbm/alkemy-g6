package models

type ProductService interface {
	GetAll() (list map[int]Product, err error)
}
